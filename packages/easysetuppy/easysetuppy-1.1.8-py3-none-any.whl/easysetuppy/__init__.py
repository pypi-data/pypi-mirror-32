import os
import sys
import shutil
import subprocess

import easygui

from easysetuppy.classifiers import clist

choices = ['name',
           'version',
           'url',
           'author',
           'author_email',
           'maintainer',
           'maintainer_email',
           'description',
           'download_url',
           'platforms',
           'license',
           ]


def none_check(val):
    """
    This checks a val

    :param val:Value to check for none
    :return: True/False
    """
    if val is None:
        easygui.msgbox(msg='Exiting program. Have a nice day!')
        sys.exit(0)


def build_and_upload(standard=False):
    """This builds and uploads a setup.py file"""
    cmds = ['build',
            'build_py',
            'build_ext',
            'build_clib',
            'build_scripts',
            'clean',
            'sdist',
            'check',
            'bdist_wheel',
            'bdist_egg',
            ]
    setupscripts = 'python setup.py '
    for entry in cmds:
        setupscripts += entry + ' '
    if os.path.exists('dist'):
        if os.listdir('dist'):
            shutil.rmtree('dist')
    subprocess.run(setupscripts)
    if not standard:
        up = easygui.multpasswordbox('Enter username and password to upload tp PyPi.',
                                     title='Username & Password',
                                     fields=['username', 'password'])
        none_check(up)
    else:
        with open('login.config', 'r') as file:
            up = eval(file.read())
    if not os.path.isfile('login.config'):
        if easygui.ccbox('Save username and password for future use?', title='Save'):
            with open('login.config', 'w') as file:
                file.write(str(up))
    username, password = up
    cmd2 = 'python -m twine upload -u {} -p {} dist\\*'.format(username, password)
    subprocess.run(cmd2)


def main():
    """The main function that makes the dialogs work."""
    global package
    di = {}
    box = easygui.multenterbox(msg='Fill in as many fields as possible. The first 3 are required.',
                               title='Setup.py Creation',
                               fields=choices)
    none_check(box)
    for i in range(len(choices)):
        if i != 1:
            di[choices[i]] = box[i]
    version_string = 'version = \'{}\''.format(box[1])
    classifiers_chosen = easygui.multchoicebox(msg='Pick the classifiers you want to use',
                                               title='Classifiers',
                                               choices=clist)
    none_check(classifiers_chosen)
    classifiers_list = classifiers_chosen
    classifiers_str = str(classifiers_chosen)
    classifiers_string = 'classifiers = ' + classifiers_str
    install = easygui.textbox(msg='Enter in packages along with their version info. Place one per line.',
                              title='Packages')
    none_check(install)
    # noinspection PyTypeChecker
    packages = install.split('\n')
    package_list = packages
    package_str = str(packages)
    package = 'install_requires = ' + package_str
    write = """from setuptools import setup, find_packages
packages = find_packages()    
{}    
{}
{}
setup(""".format(classifiers_string, package, version_string)
    for key in di:
        if not di[key]:
            pass
        else:
            write += '{}=\'{}\','.format(key, di[key])
            if key == 'name':
                write += 'version = version,'

    write += 'packages = packages'
    if not (len(package_list) == 1 and package_list[0] == ''):
        write += ', requirements = requirements'
    if not (len(classifiers_list) == 1 and classifiers_list[0] == ''):
        write += ', classifiers = classifiers'
    write += ')'
    if os.path.exists('setup.py'):
        with open('setup.py', 'r') as file:
            with open('setup.py.old', 'w') as file2:
                file2.write(file.read())
    with open('setup.py', 'w') as file:
        file.write(write)
    try:
        with open('.gitignore', 'r') as file:
            pass
        with open('.gitignore', 'a+') as file:
            if 'setup.py.old' in file.read():
                pass
            else:
                file.write('\n#Old setup.py files\nsetup.py.old')
    except FileNotFoundError:
        pass
    if easygui.ynbox(msg='Build/upload setup.py?', title='Build/Upload'):
        build_and_upload()
    return 0
