#!/bin/bash

# Functions

function programError () {
   echo "*** Program error: $1"
   exit 11
}


function quitProcess () {
   echo ""
   echo "$1"
   echo ""
   echo "Quitting ..."
   echo ""
   exit 1
}


function runCmd () {
   CMD="$1"
   echo "Running: ${CMD}"
   /bin/bash -c "${CMD}"
   result="$?"
   if [ "${result}" != 0 ]; then quitProcess "Error during running of command"; exit 1; fi
}


function showUsage () {
   echo ""
   echo "Usage of arguments: <target> <input>"
   echo ""
   echo "Where <target> is 'linux' or 'wasm'"
   echo "And <input> is a go-file or go-package specification"
   echo "The resulting file has extension .bin or .wasm, dependent of the compiler action"
   echo ""
   echo "Examples:"
   echo "   Compiling server.go to server.bin:"
   echo "      comp.sh   linux_plain   server.go"
   echo ""
   echo "Warnings:"
   echo "   1. Old versions of targets with the same name will be replaced."
   echo "   2. In case of wasm-targets: also the file 'wasm_exec.js' will be updated."     
   echo ""
   exit 1
}


function showVersions () {
   source /etc/os-release
   echo "Versions:"
   echo "   OS     : ${PRETTY_NAME}"
   echo "   Go     : $(go version)"
   echo ""
}



THIS_SCRIPT="$0"

NR_OF_ARGS="$#"
ARG1="$1"
ARG2="$2"


showVersions


# Check input parameters
if [ "${NR_OF_ARGS}" != "2" ]
then
   showUsage ""
   exit 1
fi


# Prepare compiler settings
COMPILER="${ARG1}"
SOURCE="${ARG2}"
SOURCE_BASE="$(basename ${ARG2} .go)"

echo "COMPILER = ${COMPILER}"
echo "SOURCE   = ${SOURCE}"


function forceWasmExecJsExisting () {
   if [ -f "$(go env GOROOT)/misc/wasm/wasm_exec.js" ]
   then
      # echo "OK: Found wasm_exec.js"
      return 0
   fi

   echo "Not found wasm_exec.js: Install go-misc ..."
   yum install -y golang-misc.noarch

   if [ -f "$(go env GOROOT)/misc/wasm/wasm_exec.js" ]
   then
      echo "   ==> OK: Found wasm_exec.js"
      return 0
   fi

   echo "   ==> NOK: Install failed, abort"
   exit 1
}



# Check if source exists
if [ -f "${SOURCE}" ]
then
   # Source is a file; check if it is a go-file
   if [ "${SOURCE_BASE}.go" != "${SOURCE}" ]
   then
      echo "Source is not a go-file with .go extension"
      exit 1
   fi

elif [ ! -d "src/${SOURCE}" ]
then
   echo "Source-file or Go-package not found"
   exit 1
fi


# Check compiler
if [ "${COMPILER}" == "linux" ]
then
   TARGET="${SOURCE_BASE}.bin"
   echo "TARGET   = ${TARGET}"

   ENV="GOPATH=$(go env GOPATH):$PWD CGO_ENABLED=0 GOOS=linux GOARCH=amd64"
   FLAGS='-ldflags "-w -s"'
   runCmd "${ENV} go build ${FLAGS} -o ${TARGET} ${SOURCE}"
   exit 0
fi


if [ "${COMPILER}" == "wasm" ]
then
   forceWasmExecJsExisting
   TARGET="${SOURCE_BASE}.wasm"
   echo "TARGET   = ${TARGET}"

   ENV="GOPATH=$(go env GOPATH):$PWD CGO_ENABLED=0 GOOS=js GOARCH=wasm"
   FLAGS=""

   runCmd "${ENV} go build ${FLAGS} -o ${TARGET} ${SOURCE}"
   cp $(go env GOROOT)/misc/wasm/wasm_exec.js .

   if [ ! -f "./favicon.ico" ]
   then
      echo "Not found favicon.ico; try to get the one from Golang"
      if [ -f "$(go env GOROOT)/favicon.ico" ]
      then
         cp $(go env GOROOT)/favicon.ico .
      fi
   fi

   if [ ! -f "./favicon.ico" ]
   then
      echo "WARNING: favicon.ico not found"
   fi

   exit 0
fi

showUsage "Incorrect action specified"
exit 1


