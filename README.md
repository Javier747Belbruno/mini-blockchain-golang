# Mini-Blockchain in Golang

## Intro

- Original Repo: https://github.com/ndaysinaiK/baby-blockchain
- Article: https://medium.com/swlh/is-it-hard-to-build-a-blockchain-from-scratch-2662e9b873b7

## Folder Structure

- mini-blockchain
  - AddNet - Persistence : add decentralization, fault tolerance and data persistence
  - AddPoS : Proof of Stake concensus algorithm
  - Simple : First and Simple implementation (No data persistence: Data disapears after reload)

## Prerequisites

- go installed
- nodejs installed

## Build it

### Front End

- cd mini-blockchain/AddNet-Persistence/frontend
- npm i

### Back End

- cd mini-blockchain/AddNet-Persistence/backend
- go build

## Run it

For running nodes, instructions at README.md inside AddNet-Persistence Folder.

## Credits

- Simple implementation and Article respective credits to Sinai Nday.
- AddNet-Persistence implementation credits to Javier Belbruno.
