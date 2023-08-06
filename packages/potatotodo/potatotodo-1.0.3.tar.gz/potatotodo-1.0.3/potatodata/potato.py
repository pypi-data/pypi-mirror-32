import sqlite3
import click

from db import DB
from make import Make
from remove import Remove
from modify import Modify
from find import Find
from detail import Detail
from plan import Plan
from version import Version

@click.command()

#Basic options
@click.option('--mk', nargs=3, type=str, help='Make a new plan: [descr.] [due] [category]')
@click.option('--rm', type=click.IntRange(1,), help='Remove your plan: [number]')
@click.option('--mod', help='Modify your plan: [number]')
@click.option('--find', type=str, help='Find your plan: [text]')
@click.option('--det', type=click.IntRange(1,), help='Show details of plan: [number]')

#Printing options
@click.option('--pg', type=click.IntRange(0,), default=0, help='Print the page of which you enter: [page]')
@click.option('--cat', 'p_opt', flag_value='category', help='Print the plans for category which you will select')
@click.option('--uf', 'p_opt', flag_value='unfinished', help='Print your unfinished plans')
@click.option('--f', 'p_opt', flag_value='finished', help='Print your finished plans')
@click.option('--od', 'p_opt', flag_value='overdue', help='Print your overdue plans')
@click.option('--version', is_flag = True, help='Print version')

def run(mk, rm, mod, find, det, pg, version, p_opt):
	db = DB()
	cliOption = None
	printOption = None

	#Check which option is given
	if mk:
		cliOption = Make(db, mk)
	elif rm:
		cliOption = Remove(db, rm)
	elif mod:
		cliOption = Modify(db, mod)
	elif find:
		cliOption = Find(db, find)
	elif det:
		cliOption = Detail(db, det)
	elif version:
		cliOption = Version()
	elif p_opt:
		printOption = p_opt

	if cliOption != None:
		if cliOption.check():
			cliOption.execute()
		db.conn.close()
		return

	DB().renew()
		
	Plan(db, pg, printOption).show()
	db.conn.close()

if __name__ == '__main__':
	run()




