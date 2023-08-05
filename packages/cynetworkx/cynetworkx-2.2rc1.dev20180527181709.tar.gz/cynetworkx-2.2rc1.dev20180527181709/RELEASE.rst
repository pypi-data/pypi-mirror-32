How to make a new release of ``cynetworkx``
=========================================

- Update the release notes:

  1. Review and cleanup ``doc/release/release_dev.rst``,

  2. Fix code in documentation by running
     ``cd doc && make doctest``.

  3. Make a list of merges and contributors by running
     ``doc/release/contribs.py <tag of previous release>``.

  4. Paste this list at the end of the ``release_dev.rst``. Scan the PR titles
     for highlights, deprecations, and API changes, and mention these in the
     relevant sections of the notes.

  5. Rename to ``doc/release/release_<major>.<minor>.rst``.

  6. Copy ``doc/release/release_template.rst`` to
     ``doc/release/release_dev.rst`` for the next release.

  7. Update ``doc/news.rst``.

- Toggle ``dev = True`` to ``dev = False`` in ``cynetworkx/release.py``.

- Commit changes.

- Add the version number as a tag in git::

   git tag -s [-u <key-id>] cynetworkx-<major>.<minor>

  (If you do not have a gpg key, use -m instead; it is important for
  Debian packaging that the tags are annotated)

- Push the new meta-data to github::

   git push --tags upstream master

  (where ``upstream`` is the name of the
   ``github.com:pattern-inc/cynetworkx`` repository.)

- Review the github release page::

  https://github.com/pattern-inc/cynetworkx/releases

- Publish on PyPi::

   git clean -fxd
   python setup.py sdist --formats=zip
   twine upload -s dist/cynetworkx*.zip

- Update documentation on the web:
  The documentation is kept in a separate repo: cynetworkx/documentation

  - Wait for the cyNetworkX Travis Bot to deploy to GitHub Pages
  - Sync your branch with the remote repo: ``git pull``.
  - Copy the documentation built by Travis.
    Assuming you are at the top-level of the ``documentation`` repo::

      cp -a latest cynetworkx-<major>.<minor>
      git add cynetworkx-<major>.<minor>
      ln -sfn cynetworkx-<major>.<minor> stable
      git commit -m "Add <major>.<minor> docs"
      # maybe squash all the Deploy GitHub Pages commits
      # git rebase -i HEAD~XX where XX is the number of commits back
      # check you didn't break anything
      # diff -r latest cynetworkx-<major>.<minor>
      # you will then need to force the push so be careful!
      git push

 - Increase the version number

  - Toggle ``dev = False`` to ``dev = True`` in ``cynetworkx/release.py``.
  - Update ``major`` and ``minor`` in ``cynetworkx/release.py``.

- Update the web frontpage:
  The webpage is kept in a separate repo: cynetworkx/website

  - Sync your branch with the remote repo: ``git pull``.
    If you try to ``make github`` when your branch is out of sync, it
    creates headaches.
  - Update ``_templates/sidebar_versions.html``.
  - Edit ``_static/docversions.js`` and commit
  - Push your changes to the repo.
  - Deploy using ``make github``.

