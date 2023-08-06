A minimal IPP server
====================


This is a small python script which __pretends to be a printer__.

It works well enough to print documents on my Linux box with CUPS 1.7.5, but it doesn't implement the entire IPP specification.


Add ipp-server as a printer
---------------------------

Start running the server:
```
python -m ippserver --port 1234 save /tmp/
```

The server listens to `localhost` by default. The well-known port of 631 is likely to already be in use by the web interface of CUPS, so I'm using port 1234 for this example.

Next, add the printer as you would normally on your computer (ie: the Gnome or KDE add printer dialogs). The printer location is `ipp://localhost:1234/`.


Doing things with print jobs
----------------------------

You can save print jobs as randomly named `.ps` files in a given directory:
```
python -m ippserver --port 1234 save /tmp/
```

Alternatively, you can send the postscript files to a command. The following command will run [hexdump(1)] for every print job received. hexdump reads the .ps file from stdin.
```
python -m ippserver --port 1234 run hexdump
```

Or why not email print jobs to yourself using [mail(1)]:
```
python -m ippserver --port 1234 run \
	mail -E \
	-a "MIME-Version 1.0" -a "Content-Type: application/postscript" \
	-s 'A printed document' some.person@example.com
```


PDF files
---------

The printer normally advertises itself as a postscript printer. Alternatively, the printer can advertise itself as a PDF printer. This changes the printer description (PPD), so you will need to re-add the printer (eg: with a different port).

Run the printer with `save --pdf`, and add a new printer:
```
python -m ippserver --port 7777 save --pdf /tmp/
```


[hexdump(1)]: https://linux.die.net/man/1/hexdump
[mail(1)]:  https://linux.die.net/man/1/mail
