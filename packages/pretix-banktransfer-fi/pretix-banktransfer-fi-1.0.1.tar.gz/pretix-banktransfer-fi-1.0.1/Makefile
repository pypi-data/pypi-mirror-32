all: localecompile
LNGS:=`find banktransfer_fi/locale/ -mindepth 1 -maxdepth 1 -type d -printf "-l %f "`

localecompile:
	django-admin compilemessages

localegen:
	django-admin makemessages -i build -i dist -i "*egg*" $(LNGS)

