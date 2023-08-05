To generate python documentation:

```
sudo pip install sphinx
cd docs
make html
# update the version that is checked in
rm -rf ../doc
mv build/html ../doc
# validate differences, and then
git commit
```

The doc can now be found in docs/build/html/index.html

Also note examples in `samples.py`

To run unit tests, start eeb locally:

```
./targets.sh -d update_wads
./targets.sh -d eureqa_application_server
cd Debug/eureqa_application_server
./eureqa_application_server  --js_env=development
```

And then run the tests:

```
./run_tests.sh
```


To run an individual test, do something like:

```
cd tests
python -m unittest -v test_analysis_template_runner
```

where test_analysis_template_runner is the name of a module
