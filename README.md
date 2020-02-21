# Hashcode 2020

## Introduction

This is our solution for the Google Hashcode 2020. We did not have time to submit on time but we still found a proper solution and optimized it afterwards. Now it's pretty efficient and takes < 10ms for each library to sort on my MacBook Pro.

## Install

This project requires Go to be built.

```
git clone https://github.com/snwfdhmp/hashcode-2020
cd hashcode-2020
go build -o hashcode-bin main.go
```

To scan 

```
./hashcode-bin --input a
```

## Profiling

Profiling has been made and is available in `profile.prof`. This profile is the last one after optimizations.

## Team

Project realised with @hugoglt and @eliasdemnati.