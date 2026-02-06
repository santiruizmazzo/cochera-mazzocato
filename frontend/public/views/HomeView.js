import AbstractView from "./AbstractView.js";
import ContentSection from "../components/ContentSection.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  render() {
    const homeSection = new ContentSection();
    homeSection.title = "Bienvenido a la app Cochera Mazzocato!";
    return homeSection;
  }

  async getHtml() {
    return /*html*/ `
      <section class="home-view">
        <custom-modal title="Reservar plaza">
          <slot-form></slot-form>
        </custom-modal>
  
        <header class="section-header">
          <h2>¡Bienvenido!</h2>
        </header>
        
        <slots-collection></slots-collection>
        <div></div>
      </section>
    `;
  }

  setUpJavascript() {
    const view = document.querySelector(".home-view");
    const modal = document.querySelector("custom-modal");
    const form = document.querySelector("slot-form");

    view.addEventListener("slot:selected", (event) => {
      form.slotId = event.detail.slotId;
      form.slotNumber = event.detail.slotNumber;
      modal.show();
    });

    view.addEventListener("slot:assigned", () => modal.close());
  }
}
