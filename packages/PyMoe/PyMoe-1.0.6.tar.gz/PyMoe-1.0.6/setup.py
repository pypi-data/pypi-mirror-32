from setuptools import setup

setup(
    name='PyMoe',
    version='1.0.6',
    packages=['Pymoe', 'Pymoe.Anilist', 'Pymoe.Kitsu', 'Pymoe.VNDB', 'Pymoe.Bakatsuki'],
    url='https://github.com/ccubed/PyMoe',
    license='MIT',
    author='Cooper Click',
    author_email='ccubed.techno@gmail.com',
    description="PyMoe is the only lib you'll ever need if you need the animu or mangu on the Python Platform. It supports AniList, VNDB, Kitsu and AniDB.",
    classifiers=[
        'Development Status :: 5 - Production/Stable',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3.6',
        'Programming Language :: Python :: 3.5',
        'Topic :: Software Development :: Libraries'
    ],
    keywords="Anime Manga LN VN VNDB Anilist Kitsu AniDB MyAnimeList MAL Bakatsuki",
    install_requires=['requests', 'bs4', 'ujson']
)
