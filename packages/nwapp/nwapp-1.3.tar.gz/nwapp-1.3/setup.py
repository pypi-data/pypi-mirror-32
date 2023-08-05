


from setuptools import setup, find_packages


setup(
        name = 'nwapp',
        version = '1.3',
        description='This CLI tool is to manage freelance works',
        url='http://www.example.com/~cschultz/bvote/',
        author='Naseef Ummer',
        author_email = 'naseefo@gmail.com',
        license = 'MIT',
        zip_safe = False,
        #py_modules = ['hello'],
        packages = find_packages(),
        include_package_data=True,
        install_requires = [
            'Click',
            'pymongo',
            'pandas'
            ],
        entry_points = {
            'console_scripts':[
                'nw = nwapp.scripts.nw:cli',

            ],
        }
    )
