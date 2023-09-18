export const initWasm = async () => {
    const go = new Go();
    const { instance } = await WebAssembly.instantiateStreaming(
        fetch("../../wasm/static/main.wasm"),
        go.importObject,
    )
    go.run(instance)
    const add = instance.exports.add;
    return { add }
}