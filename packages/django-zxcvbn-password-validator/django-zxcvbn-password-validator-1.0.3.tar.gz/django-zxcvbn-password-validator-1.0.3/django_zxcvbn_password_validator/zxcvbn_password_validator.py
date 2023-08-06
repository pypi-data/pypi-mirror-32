from django.conf import settings
from django.core.exceptions import ValidationError, ImproperlyConfigured
from django.utils.translation import ugettext_lazy as _
from zxcvbn import zxcvbn


class ZxcvbnPasswordValidator(object):

    def __init__(self, min_length=1):
        self.min_length = min_length
        error_msg = "ZxcvbnPasswordValidator need an integer between 0 and 4 "
        error_msg += "for PASSWORD_MINIMAL_STRENGH in the settings."
        try:
            self.password_minimal_strengh = settings.PASSWORD_MINIMAL_STRENGH
        except AttributeError:
            raise ImproperlyConfigured(error_msg)
        if int(self.password_minimal_strengh) != self.password_minimal_strengh:
            error_msg += f" (not a {self.password_minimal_strengh.__class__.__name__})"
            raise ImproperlyConfigured(error_msg)

    def validate(self, password, user=None):
        results = zxcvbn(password, user_inputs=user.__dict__)
        password_strengh = results["score"]
        if password_strengh < self.password_minimal_strengh:
            crack_time = results["crack_times_display"]
            offline_time = crack_time["offline_slow_hashing_1e4_per_second"]
            warn = results["feedback"]["warning"]
            advice = results["feedback"]["suggestions"]
            comment = "{} {}".format(
                _(f'Your password is too guessable :'),
                _(f'It would take an offline attacker {offline_time} to guess it.'),
            )
            if warn:
                comment += " {}".format(_(f'Warning : {warn}.'))
            if advice:
                comment += " {}".format(_(f'What you can do : {advice}.'))
            raise ValidationError(comment)

    def get_help_text(self):
        expectations = _("We expect")
        hardness = {
            0: _("a password that is not one of the 1000 most used worldwide."),
            1: _("a password that cannot be guessed by slow online attacks."),
            2: _("a password that cannot be guessed by fast online attacks."),
            3: _("a password that could not be guessed if a lone engineer hacked into our database."),
            4: _("a password that could not be guessed if a large organisation hacked into our database."),
        }
        expectations += " {}".format(hardness.get(
            self.password_minimal_strengh,
            _("our developpers to change the value for 'PASSWORD_MINIMAL_STRENGH' to something valid ASAP, sorry :)")
        ))
        return "{} {} {} {}".format(
            _("There is no specific rule for a great password,"),
            _("however if your password is too easy to guess,"),
            _("we will tell you how to make a better one."),
            expectations
        )
