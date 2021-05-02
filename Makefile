PY?=python3
WORKDIR?=.
VENVDIR=$(WORKDIR)/env
VENV=$(VENVDIR)/bin

.PHONY: run
run: $(VENV)/.marker
	. env/bin/activate && python3 -m vidhasher $(ARGS)

.PHONY: venv
venv: $(VENV)/.marker

$(VENV)/:
	$(PY) -m venv $(VENVDIR)
	$(VENV)/python -m pip install --upgrade pip setuptools wheel

$(VENV)/.marker: setup.py
	$(VENV)/pip install -e .
	touch $(VENV)/.marker
