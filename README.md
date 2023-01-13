# Chain link tool

Chain link tool allows to create the chain link proof for linking Desmos profile.

## Installation

To install `chain-link-tool`, run the following command:
```bash
git clone https://github.com/desmos-labs/chain-link-tool.git
cd chain-link-tool
make install
```

## Usage

### Create chain link json

In order to generate the needed chain link JSON ,you can run this command for the target provider chain:
```bash
chain-link-tool generate [provider]
```

For example:
```bash
chain-link-tool generate nomic
```

Once you have built the JSON object using this command, you can then run the following command to complete the linkage:
```
desmos tx profiles link-chain [/path/to/json/file.json]
```

### Support providers

Currently, chain link tool supports following chains:
* Nomic