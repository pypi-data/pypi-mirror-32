# django-zxcvbn-password-validator

A password validator for django, based on zxcvbn-python and available with pip.

[![PyPI version](https://badge.fury.io/py/django-zxcvbn-password-validator.svg)](https://badge.fury.io/py/django-zxcvbn-password-validator)

# How to use

Add it to your requirements and get it with pip.

	django-zxcvbn-password-validator==1.0.2

Then everything happens in your settings file.

Add `'django_zxcvbn_password_validator'` in the `INSTALLED_APPS` :

	INSTALLED_APPS = [
		...
		'django_zxcvbn_password_validator'
	]

Modify `AUTH_PASSWORD_VALIDATORS` :

	AUTH_PASSWORD_VALIDATORS = [
		{
			'NAME': 'django_zxcvbn_password_validator.ZxcvbnPasswordValidator',
		},
		...
	]

You could choose to use zxcvbn alone, but I personally still use Django's `UserAttributeSimilarityValidator`.

Finally set the `PASSWORD_MINIMAL_STRENGH` to your liking, every password scoring
lower than this number will be rejected :

	# 0 too guessable: risky password. (guesses < 10^3)
	# 1 very guessable: protection from throttled online attacks. (guesses < 10^6)
	# 2 somewhat guessable: protection from unthrottled online attacks. (guesses < 10^8)
	# 3 safely unguessable: moderate protection from offline slow-hash scenario. (guesses < 10^10)
	# 4 very unguessable: strong protection from offline slow-hash scenario. (guesses >= 10^10)
	PASSWORD_MINIMAL_STRENGH = 3

