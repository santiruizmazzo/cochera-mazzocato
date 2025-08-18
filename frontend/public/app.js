async function obtenerVersion() {
  try {
    const response = await fetch(window.ENV.API_URL + "/health");
    if (!response.ok) {
      return "☹";
    }
    return await response.json().then((data) => {
      return "v" + data["version"];
    });
  } catch (error) {
    console.error("Hubo un problema con el fetch:", error);
    return "☹";
  }
}

var cajaDeVersion = document.querySelector(".caja-version");
cajaDeVersion.innerHTML = await obtenerVersion();
