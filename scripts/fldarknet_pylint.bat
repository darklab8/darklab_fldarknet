
cmd /k "cd /d ..\venv\Scripts & activate & cd /d    ..\.. & pylint --load-plugins pylint_django --django-settings-module="fldarknet.settings" --disable=django-not-configured fldarknet main ship commodities parsing"