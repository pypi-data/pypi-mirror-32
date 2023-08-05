send email with one line using ssl
you need a local_settings.py in the root of the virtualenv with this:

# MAIL_CONFIG_DICT is required by for senderror_simple which is a function of the module quickmail 
# SMTP_SSL_HOST is a ssl enabled smtp host and SMTP_SSL_PORT is the port to use (usually 465)
# MAIL_USER and MAIL_PASS refer to the email account to use to send error messages. MAIL_RECIPIENTS is a list of emails (strings) to be notified of errors
MAIL_CONFIG_DICT = {
    'SMTP_SSL_HOST': 'yourdomain.com',
    'SMTP_SSL_PORT': 465,
    'MAIL_USER': 'you@yourdomain.com',
    'MAIL_PASS': 'yourpass',
    'MAIL_RECIPIENTS': ['you@yourdomain.com', friend@anotherdomain.com']
    }
