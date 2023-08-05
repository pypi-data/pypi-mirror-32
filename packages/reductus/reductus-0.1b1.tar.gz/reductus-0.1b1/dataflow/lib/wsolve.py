r"""
Weighted linear and polynomial solver with uncertainty.

Given $A \bar x = \bar y \pm \delta \bar y$, solve using *s = wsolve(A, y, dy)*

*wsolve* uses the singular value decomposition for increased accuracy.

The uncertainty in the solution is estimated from the scatter in the data.
Estimates the uncertainty for the solution from the scatter in the data.

The returned model object *s* provides:

    ======== ============================================
    ======== ============================================
    s.x      solution
    s.std    uncertainty estimate assuming no correlation
    s.rnorm  residual norm
    s.DoF    degrees of freedom
    s.cov    covariance matrix
    s.ci(p)  confidence intervals at point p
    s.pi(p)  prediction intervals at point p
    s(p)     predicted value at point p
    ======== ============================================

**Example**

Weighted system::

    >>> import numpy as np
    >>> import wsolve
    >>> A = np.matrix("1,2,3;2,1,3;1,1,1", 'd').A
    >>> dy = [0.2, 0.01, 0.1]
    >>> y = [ 14.16, 13.01, 6.15]
    >>> s = wsolve.wsolve(A, y, dy)
    >>> print(", ".join("%0.2f +/- %0.2f"%(a, b) for a, b in zip(s.x, s.std)))
    1.05 +/- 0.17, 2.20 +/- 0.12, 2.91 +/- 0.12


Note there is a counter-intuitive result that scaling the estimated
uncertainty in the data does not affect the computed uncertainty in
the fit.  This is the correct result --- if the data were indeed
selected from a process with ten times the uncertainty, you would
expect the scatter in the data to increase by a factor of ten as
well.  When this new data set is fitted, it will show a computed
uncertainty increased by the same factor.  Monte carlo simulations
bear this out.  The conclusion is that the dataset carries its own
information about the variance in the data, and the weight vector
serves only to provide relative weighting between the points.
"""

__all__ = ['wsolve', 'wpolyfit', 'LinearModel', 'PolynomialModel', 'smooth']

try:
    #from typing import Optional, Union, Sequence
    pass
except ImportError:
    pass

# FIXME: test second example
#
# Example 2: weighted overdetermined system  y = x1 + 2*x2 + 3*x3 + e
#
#    A = fullfact([3, 3, 3]); xin=[1;2;3];
#    y = A*xin; dy = rand(size(y))/50; y+=dy.*randn(size(y));
#    [x, s] = wsolve(A, y, dy);
#    dx = s.normr*sqrt(sumsq(inv(s.R'))'/s.df);
#    res = [xin, x, dx]


import numpy as np


class LinearModel(object):
    r"""
    Model evaluator for linear solution to $Ax = y$.

    Use *s(A)* to compute the predicted value of the linear model *s*
    at points given on the rows of $A$.

    Computes a confidence interval (range of likely values for the
    mean at $x$) or a prediction interval (range of likely values
    seen when measuring at $x$).  The prediction interval gives
    the width of the distribution at $x$.  This should be the same
    regardless of the number of measurements you have for the value
    at $x$.  The confidence interval gives the uncertainty in the
    mean at $x$.  It should get smaller as you increase the number of
    measurements.  Error bars in the physical sciences usually show
    a $1-\alpha$ confidence value of $\text{erfc}(1/\sqrt{2})$, representing
    a $1-\sigma$ standand deviation of uncertainty in the mean.

    Confidence intervals for the expected value of the linear system
    evaluated at a new point $w$ are given by the $t$ distribution for
    the selected interval $1-\alpha$, the solution $x$, and the number
    of degrees of freedom $n-p$:

    .. math::

        w^T x \pm t^{\alpha/2}_{n-p} \sqrt{ \text{var}(w) }

    where the variance $\text{var}(w)$ is given by:

    .. math::

        \text{var}(w) = \sigma^2 (w^T (A^TA)^{-1} w)

    Prediction intervals are similar, except the variance term increases to
    include both the uncertainty in the predicted value and the variance in
    the data:

    .. math::

        \text{var}(w) = \sigma^2 (1 + w^T (A^TA)^{-1} w)
    """
    def __init__(self, x=None, DoF=None, SVinv=None, rnorm=None):
        # type: (np.ndarray, int, np.ndarray, float) -> None
        # Note: SVinv should be computed from S, V where USV' = A
        #: solution to the equation $Ax = y$
        self.x = x
        #: number of degrees of freedom in the solution
        self.DoF = DoF
        #: 2-norm of the residuals $||y-Ax||_2$
        self.rnorm = rnorm
        self._SVinv = SVinv

    def __call__(self, A):
        """
        Return the prediction for a linear system at points in the rows of A.
        """
        return np.dot(np.asarray(A), self.x)

    # covariance matrix invC = A'A  = (USV')'USV' = VSU'USV' = VSSV'
    # C = inv(A'A) = inv(VSSV') = inv(V')inv(SS)inv(V) = Vinv(SS)V'
    # diag(inv(A'A)) is sum of the squares of the columns inv(S) V'
    # and is also the sum of the squares of the rows of V inv(S)
    @property
    def cov(self):
        """covariance matrix [inv(A'A); O(n^3)]"""
        # FIXME: don't know if we need to scale by C, but it will
        # at least make things consistent
        C = self.rnorm**2/self.DoF if self.DoF > 0 else 1
        return C * np.dot(self._SVinv, self._SVinv.T)

    @property
    def var(self):
        """solution variance [diag(cov); O(n^2)]"""
        C = self.rnorm**2/self.DoF if self.DoF > 0 else 1
        return C * np.sum(self._SVinv**2, axis=1)

    @property
    def std(self):
        """solution standard deviation [sqrt(var); O(n^2)]"""
        return np.sqrt(self.var)

    @property
    def p(self):
        """p-value probability of rejection"""
        from scipy.stats import chi2  # lazy import in case scipy not present
        return chi2.sf(self.rnorm ** 2, self.DoF)

    def _interval(self, X, alpha, pred):
        """
        Helper for computing prediction/confidence intervals.
        """
        # Comments from QR decomposition solution to Ax = y:
        #
        #   Rather than A'A we have R from the QR decomposition of A, but
        #   R'R equals A'A.  Note that R is not upper triangular since we
        #   have already multiplied it by the permutation matrix, but it
        #   is invertible.  Rather than forming the product R'R which is
        #   ill-conditioned, we can rewrite x' inv(A'A) x as the equivalent
        #      x' inv(R) inv(R') x = t t', for t = x' inv(R)
        #
        # We have since switched to an SVD solver, which gives us
        #
        #    invC = A'A  = (USV')'USV' = VSU'USV' = VSSV'
        #    C = inv(A'A) = inv(VSSV') = inv(V')inv(SS)inv(V)
        #      = Vinv(SS)V' = Vinv(S) inv(S)V'
        #
        # Substituting, we get
        #
        #    x' inv(A'A) x = t t', for t = x' Vinv(S)
        #
        # Since x is a vector, t t' is the inner product sum(t**2).
        # Note that LAPACK allows us to do this simultaneously for many
        # different x using sqrt(sum(T**2, axis=1)), with T = X' Vinv(S).
        #
        # Note: sqrt(F(1-a;1, df)) = T(1-a/2;df)
        #
        from scipy.stats import t  # lazy import in case scipy not present
        y = np.dot(X, self.x).ravel()
        s = t.ppf(1-alpha/2, self.DoF) * self.rnorm/np.sqrt(self.DoF)
        t = np.dot(X, self._SVinv)
        dy = s * np.sqrt(pred + np.sum(t**2, axis=1))
        return y, dy

    def ci(self, A, sigma=1):
        r"""
        Compute the calculated values and the confidence intervals
        for the linear model evaluated at $A$.

        *sigma=1* corresponds to a $1-\sigma$ confidence interval

        Confidence intervals are sometimes expressed as $1-\alpha$ values,
        where $\alpha = \text{erfc}(\sigma/\sqrt{2})$.
        """
        from scipy.special import erfc  # lazy import in case scipy not present
        alpha = erfc(sigma / np.sqrt(2))
        return self._interval(np.asarray(A), alpha, 0)

    def pi(self, A, p=0.05):
        r"""
        Compute the calculated values and the prediction intervals
        for the linear model evaluated at $A$.

        *p=0.05* corresponds to the 95% prediction interval.
        """
        return self._interval(np.asarray(A), p, 1)


def wsolve(A, y, dy=1.0, rcond=1e-12):
    # type: (Sequence[Sequence[float]], Sequence[float], Sequence[float], float) -> LinearModel
    r"""
    Given a linear system $y = A x + \delta y$, estimates $x$ and $\delta x$.

    *A* is an n x m array of measurement points.

    *y* is an n x k array or vector of length n of measured values at *A*.

    *dy* is a scalar or an n x 1 array of uncertainties in the values at *A*.

    Returns :class:`LinearModel`.
    """
    # The ugliness v[:, N.newaxis] transposes a vector
    # The ugliness N.dot(a, b) is a*b for a, b matrices
    # The ugliness vh.T.conj() is the hermitian transpose

    # Make sure inputs are arrays
    A, y, dy = np.asarray(A), np.asarray(y), np.asarray(dy)
    if dy.ndim == 1:
        dy = dy[:, np.newaxis]
    if y.ndim == 1:
        y = y[:, np.newaxis]

    # Apply weighting if dy is not a scalar
    # If dy is a scalar, it cancels out of both sides of the equation
    # Note: with A, dy arrays instead of matrices, A/dy operates element-wise
    # Since dy is a row vector, this divides each row of A by the corresponding
    # element of dy.
    if dy.ndim == 2:
        A, y = A/dy, y/dy

    # Singular value decomposition: A = U S V.H
    # Since A is an array, U, S, VH are also arrays
    # The zero indicates an economy decomposition, with u nxm rathern than nxn
    u, s, vh = np.linalg.svd(A, full_matrices=False)

    # FIXME what to do with ill-conditioned systems?
    #if s[-1]<rcond*s[0]: raise ValueError, "matrix is singular"
    # s[s<rcond*s[0]] = 0.  # Can't do this because 1/s below will fail

    # Solve: x = V inv(S) U.H y
    # S diagonal elements => 1/S is inv(S)
    # A*D, D diagonal multiplies each column of A by the corresponding diagonal
    # D*A, D diagonal multiplies each row of A by the corresponding diagonal
    # Computing V*inv(S) is slightly faster than inv(S)*U.H since V is smaller
    # than U.H.  Similarly, U.H*y is somewhat faster than V*U.H
    SVinv = vh.T.conj() / s
    Uy = np.dot(u.T.conj(), y)
    x = np.dot(SVinv, Uy)

    DoF = y.shape[0] - x.shape[0]
    rnorm = np.linalg.norm(y - np.dot(A, x))

    return LinearModel(x=x, DoF=DoF, SVinv=SVinv, rnorm=rnorm)


def _poly_matrix(x, degree, origin=False):
    # type: (Sequence[float], int, bool) -> np.ndarray
    """
    Generate the matrix A used to fit a polynomial using a linear solver.
    """
    if origin:
        n = np.arange(degree, 0, -1, dtype='i')
    else:
        n = np.arange(degree, -1, -1, dtype='i')
    x = np.asarray(x, 'd')
    return x[:, None] ** n[None, :]


class PolynomialModel(object):
    r"""
    Model evaluator for best fit polynomial $p(x) = y +/- \delta y$.

    Use *p(x)* for PolynomialModel *p* to evaluate the polynomial at all
    points in the vector *x*.
    """

    def __init__(self, s, origin=False, data=None):
        #: True if polynomial goes through the origin
        self.data = data
        self.origin = origin
        #: polynomial coefficients
        self.coeff = np.ravel(s.x)
        if origin:
            self.coeff = np.hstack((self.coeff, 0))
        #: polynomial degree
        self.degree = len(self.coeff) - 1
        #: number of degrees of freedom in the solution
        self.DoF = s.DoF
        #: 2-norm of the residuals $||y-Ax||_2$
        self.rnorm = s.rnorm
        self._conf = s

    @property
    def cov(self):
        """
        covariance matrix

        Note that the ones column will be absent if *origin* is True.
        """
        return self._conf.cov

    @property
    def var(self):
        """solution variance"""
        return np.hstack((self._conf.var, 0)) if self.origin else self._conf.var

    @property
    def std(self):
        """solution standard deviation"""
        return np.hstack((self._conf.std, 0)) if self.origin else self._conf.std

    @property
    def p(self):
        """p-value probability of rejection"""
        return self._conf.p

    def __call__(self, x):
        """
        Evaluate the polynomial at x.
        """
        return np.polyval(self.coeff, x)

    def der(self, x):
        """
        Evaluate the polynomial derivative at x.
        """
        return np.polyval(np.polyder(self.coeff), x)

    def ci(self, x, sigma=1):
        """
        Evaluate the polynomial and the confidence intervals at x.

        sigma=1 corresponds to a 1-sigma confidence interval
        """
        A = _poly_matrix(x, self.degree, self.origin)
        return self._conf.ci(A, sigma)

    def pi(self, x, p=0.05):
        """
        Evaluate the polynomial and the prediction intervals at x.

        p = 1-alpha = 0.05 corresponds to 95% prediction interval
        """
        A = _poly_matrix(x, self.degree, self.origin)
        return self._conf.pi(A, p)

    def __str__(self):
        # TODO: better polynomial pretty printing using formatnum
        return "Polynomial(%s)" % self.coeff


def wpolyfit(x, y, dy=1.0, degree=None, origin=False):
    # type: (Sequence[float], Sequence[float], Union[Sequence[float], float], Optional[int], bool) -> PolynomialModel
    r"""
    Return the polynomial of degree $n$ that minimizes $\sum(p(x_i) - y_i)^2/\sigma_i^2$.

    if origin is True, the fit should go through the origin.

    Returns :class:`PolynomialModel`.
    """
    assert degree is not None, "Missing degree argument to wpolyfit"

    A = _poly_matrix(x, degree, origin)
    s = wsolve(A, y, dy)
    return PolynomialModel(s, origin=origin, data=(x, y, dy))


def wpolyplot(poly, with_pi=False, with_ci=True):
    # type: (PolynomialModel, bool, bool) -> None
    import pylab
    x, y, dy = poly.data
    pylab.errorbar(x, y, yerr=dy, fmt='gx')
    pylab.hold(True)
    px = np.linspace(x.min(), x.max(), 400)
    py, pdy = poly.pi(px)
    cy, cdy = poly.ci(px)
    pylab.plot(px, py, 'g-')
    if with_pi:
        pylab.plot(px, py + pdy, 'g-.', px, py - pdy, 'g-.', hold=True)
    if with_ci:
        pylab.plot(px, cy + cdy, 'r-.', px, cy - cdy, 'r-.', hold=True)


def smooth(x, xp, yp, dyp=None, degree=2, span=5):
    # type: (Sequence[float], Sequence[float], Sequence[float], Sequence[float], int, int) -> (np.ndarray, np.ndarray)
    """
    Moving least squares smoother.

    *x* are the points at which to evaluate the smoothed values, returning
    estimated *y +/- dy* for each point, where *dy* is the confidence interval
    for the value of *y*.

    *xp*, *yp*, *dyp* are the original data points and uncertainties.

    *degree* is the degree of the polynomial smoother, and *span* is the
    number of neighbouring points to use. If *xp* is equally spaced and
    evaluated at *x=xp*, then this is equivalent to a Savitsky-Golay filter
    of the same degree and span.  Linear interpolation is *degree=1* and
    *span=2*, but for that it will be much faster to use a linear interpolation
    function directly.

    Note: *xp* does not need to be equally spaced, but if it has large clumps
    then a fixed window length is not the most appropriate algorithm.
    Instead, one could specify a window width *dx* for each *x*, and use
    all data points within *[x-dx, x+dx]* as the input to the polynomial.
    This capability is not yet implemented.
    """
    if dyp is None:
        dyp = np.ones(len(yp))

    if len(xp) <= span:
        n = len(xp)
        if n <= degree:
            degree = n-1
        poly = wpolyfit(xp, yp, dyp, degree=degree)
        y, dy = poly.ci(x)

    else:
        y, dy = np.empty(x.shape), np.empty(x.shape)
        if span%2 == 0:
            # Even span is an odd number of intervals, so set boundaries
            # at the x points.
            index = np.searchsorted(xp[span//2:-span//2], x)
        else:
            # Odd span is even number of intervals, so set boundaries
            # at the midpoints between x.
            index = np.searchsorted(0.5*(xp[:-span]+xp[span:]), x)

        # Save a little when densely interpolating by caching poly
        # Since indices returned by searchsorted are always non-negative,
        # initializing last to -1 guarantees that the first time through the
        # loop last != start, and a new polynomial will be fit.
        last = -1
        for i, start in enumerate(index):
            if last != start:
                # Note that the centers are offset by -span//2 because the
                # search started that far into the xp array.
                s = slice(start, start+span)
                poly = wpolyfit(xp[s], yp[s], dyp[s], degree=degree)
            #if i%10 == 0:
            #    from ..wsolve import wpolyplot;wpolyplot(poly)
            y[i], dy[i] = [v[0] for v in poly.ci([y[i]])]

    return y, dy


def demo():
    """
    Fit a random cubic polynomial.
    """
    # Make fake data
    x = np.linspace(-15, 5, 15)
    th = np.polyval([.2, 3, 1, 5], x)  # polynomial
    dy = np.sqrt(np.abs(th))        # poisson uncertainty estimate
    y = np.random.normal(th, dy)    # but normal generator

    # Fit to a polynomial
    poly = wpolyfit(x, y, dy=dy, degree=3)

    wpolyplot(poly, with_pi=True, with_ci=True)
    import pylab; pylab.show()

def demo2():
    x = [1, 2, 3]
    y = [10, 8, 6]
    dy = [1, 3, 1]
    poly = wpolyfit(x, y, dy=dy, degree=1)

    wpolyplot(poly)
    import pylab; pylab.show()


def test():
    """
    Check that results are correct for a known problem.
    """
    x = np.array([0, 1, 2, 3, 4], 'd')
    y = np.array([2.5, 7.9, 13.9, 21.1, 44.4], 'd')
    dy = np.array([1.7, 2.4, 3.6, 4.8, 6.2], 'd')
    poly = wpolyfit(x, y, dy, 1)
    px = np.array([1.5], 'd')
    _, pi = poly.pi(px)  # Same y is returend from pi and ci
    py, ci = poly.ci(px)

    # Uncomment these to show target values
    # print "Tp = [%.16g, %.16g]"%(p[0], p[1])
    # print "Tdp = [%.16g, %.16g]"%(dp[0], dp[1])
    # print "Tpi, Tci = %.16g, %.16g"%(pi, ci)
    Tp = np.array([7.787249069840737, 1.503992847461524])
    Tdp = np.array([1.522338103010216, 2.117633626902384])
    Tpi, Tci = 7.611128464981324, 2.342860389884832

    perr = np.max(np.abs(poly.coeff - Tp))
    dperr = np.max(np.abs(poly.std - Tdp))
    cierr = np.abs(ci - Tci)
    pierr = np.abs(pi - Tpi)
    assert perr < 1e-14, "||p-Tp||=%g" % perr
    assert dperr < 1e-14, "||dp-Tdp||=%g" % dperr
    assert cierr < 1e-14, "||ci-Tci||=%g" % cierr
    assert pierr < 1e-14, "||pi-Tpi||=%g" % pierr
    assert py == poly(px), "direct call to poly function fails"

if __name__ == "__main__":
#    test()
#    demo()
    demo2()
