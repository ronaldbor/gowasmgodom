
async function WASM_loadFiles () {
   for (let i = 0; i < arguments.length; i++) {
      await WASM_activateFile (arguments[i]);
   }
}


async function WASM_activateFile (wasmName) {
   console.log (`JS: Activating ${wasmName} ...`);

   const go = new Go ();
   WebAssembly.instantiateStreaming (
      fetch (wasmName), go.importObject)
   .then ((result) => {
      go.run (result.instance);
   });
   console.log (`JS: ... ready activating ${wasmName}`);
}

