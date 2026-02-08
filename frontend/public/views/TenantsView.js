import AbstractView from "./AbstractView.js";
import ContentSection from "../components/ContentSection.js";
import TenantsCollection from "../components/TenantsCollection.js";
import TenantForm from "../components/TenantForm.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  renderWithin(mainElement) {
    this.fetchTenants().then((tenants) => {
      this.tenantsCollection.load(tenants);
    });

    this.homeSection = new ContentSection();

    this.homeSection.title = "Inquilinos registrados";
    this.homeSection.buttonText = "Nuevo inquilino";

    this.tenantForm = new TenantForm();
    this.homeSection.modalTitle = "Registrar inquilino";
    this.homeSection.modalContent = this.tenantForm;

    this.tenantsCollection = new TenantsCollection();

    this.homeSection.content = this.tenantsCollection;

    mainElement.appendChild(this.homeSection);
  }

  async fetchTenants() {
    const TENANTS_URL = import.meta.env.VITE_API_URL + "/api/tenants";

    return await fetch(TENANTS_URL)
      .then((response) => response.json())
      .then((json) => {
        return json["data"];
      })
      .catch((error) => {
        console.error("Error fetching tenants:", error);
      });
  }

  getHtml() {
    return /*html*/ `
      <section class="tenants-view">
        <custom-modal title="Registrar inquilino">
          <tenant-form mode="create"></tenant-form>
        </custom-modal>

        <header class="section-header">
          <h2>Inquilinos registrados</h2>
          <open-modal-button>
            <svg><use href="#person_add"></svg>
            Nuevo inquilino
          </open-modal-button>
        </header>
        
        <tenants-collection></tenants-collection>
        <div></div>
      </section>
    `;
  }

  setUpJavascript() {
    const view = document.querySelector(".tenants-view");
    const modal = document.querySelector("custom-modal");

    view.addEventListener("tenants:created", () => modal.close());
  }
}
