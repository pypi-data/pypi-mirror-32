#! python3

"""Comic Crawler

Usage:
  comiccrawler [--profile=<profile>] (
    domains |
    download <url> [--dest=<save_path>] |
    gui |
    migrate
  )
  comiccrawler (--help | --version)

Commands:
  domains             List supported domains.
  download <url>      Download from the URL.
  gui                 Launch TKinter GUI.
  migrate             Convert old file path to new file path.

Options:
  --profile=<profile> Set profile location. [default: ~/comiccrawler]
  --dest=<save_path>  Set download save path. [default: .]
  --help              Show help message.
  --version           Show current version.

Sub modules:
  comiccrawler.core   Core functions of downloading, analyzing.
  comiccrawler.error  Errors.
  comiccrawler.mods   Import download modules.
"""

from .__pkginfo__ import __version__

def console_download(url, savepath):
	"""Download url to savepath."""
	from .core import Mission, download, analyze

	mission = Mission(url=url)
	analyze(mission)
	download(mission, savepath)

def console_init():
	"""Console init."""
	from docopt import docopt

	arguments = docopt(__doc__, version=__version__)
	
	if arguments["--profile"]:
		from .profile import set as set_profile
		set_profile(arguments["--profile"])

	if arguments["domains"]:
		from .mods import list_domain
		print("Supported domains:\n" + ", ".join(list_domain()))

	elif arguments["gui"]:
		from .gui import main
		main()

	elif arguments["download"]:
		console_download(arguments["<url>"], arguments["--dest"])
		
	elif arguments["migrate"]:
		migrate()

def migrate():
	import re
	
	from .mission_manager import mission_manager, get_mission_id
	from .core import safefilepath
	from .io import move
	from .safeprint import print
	from .profile import get as profile
	
	def safefilepath_old(file):
		"""Return a safe directory name."""
		return re.sub(r"[/\?|<>:\"*]", "_", file).strip()
		
	def rename(to_rename):
		for src, dst in to_rename:
			if src != dst:
				print("\n" + src + "\n" + dst)
				move(src, dst)
		
	mission_manager.load()
	to_rename = []
	
	for mission in mission_manager.pool.values():
		id = get_mission_id(mission)
		old_ep = profile("pool/" + safefilepath_old(id + ".json"))
		new_ep = profile("pool/" + safefilepath(id + ".json"))
		to_rename.append((old_ep, new_ep))
		
	rename(to_rename)
	