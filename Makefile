PY?=python3
WORKDIR?=.
VENVDIR=$(WORKDIR)/env
VENV=$(VENVDIR)/bin
SRC=vidhasher

.PHONY: run
run: $(VENV)/.marker $(SRC)/contract.bin
	. env/bin/activate && python3 -m vidhasher $(ARGS)

$(SRC)/contract.bin: $(SRC)/contract.sol
	cd $(SRC) && solc --bin contract.sol | tail -n +4 > contract.bin
$(SRC)/contract.abi: $(SRC)/contract.sol
	cd $(SRC) && solc --abi contract.sol | tail -n +4 > contract.abi

.PHONY: venv
venv: $(VENV)/.marker

$(VENV)/:
	$(PY) -m venv $(VENVDIR)
	$(VENV)/python -m pip install --upgrade pip setuptools wheel

$(VENV)/.marker: setup.py
	$(VENV)/pip install -e .
	touch $(VENV)/.marker
