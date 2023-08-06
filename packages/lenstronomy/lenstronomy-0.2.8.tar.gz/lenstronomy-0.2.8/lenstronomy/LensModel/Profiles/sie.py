from lenstronomy.LensModel.Profiles.spemd import SPEMD


class SIE(object):
    """
    class for singular isothermal ellipsoid (SIS with ellipticity)
    """
    def __init__(self):
        self.spemd = SPEMD()
        self._gamma = 2

    def function(self, x, y, theta_E, e1, e2, center_x=0, center_y=0):
        """

        :param x:
        :param y:
        :param theta_E:
        :param q:
        :param phi_G:
        :param center_x:
        :param center_y:
        :return:
        """
        return self.spemd.function(x, y, theta_E, self._gamma, e1, e2, center_x, center_y)

    def derivatives(self, x, y, theta_E, e1, e2, center_x=0, center_y=0):
        """

        :param x:
        :param y:
        :param theta_E:
        :param q:
        :param phi_G:
        :param center_x:
        :param center_y:
        :return:
        """
        return self.spemd.derivatives(x, y, theta_E, self._gamma, e1, e2, center_x, center_y)

    def hessian(self, x, y, theta_E, e1, e2, center_x=0, center_y=0):
        """

        :param x:
        :param y:
        :param theta_E:
        :param q:
        :param phi_G:
        :param center_x:
        :param center_y:
        :return:
        """
        return self.spemd.hessian(x, y, theta_E, self._gamma, e1, e2, center_x, center_y)