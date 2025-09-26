module.exports = {
  plugins: [
    "removeXMLNS",
    "removeXMLProcInst",
    "removeDoctype",
    "removeComments",
    "removeMetadata",
    "removeTitle",
    "removeDesc",
    { name: "removeViewBox", active: false },
    {
      name: "removeAttrs",
      params: {
        attrs: "(fill|stroke)",
      },
    },
    "removeUselessDefs",
    "removeDimensions",
    "cleanupAttrs",
    "convertPathData",
    "mergePaths",
    "removeEmptyAttrs",
    "removeHiddenElems",
    "removeEmptyText",
    "removeUnknownsAndDefaults",
  ],
};
