const initContactForm = () => {
  const form = document.querySelector("[data-contact-form]");
  if (!form) {
    return;
  }

  const submitButton = form.querySelector("[data-contact-submit]");
  const statusElement = form.querySelector("[data-contact-status]");
  const apiBase = document.body.dataset.apiBase || window.location.origin;
  const submitUrl = new URL("/api/contact", apiBase).toString();

  const setStatus = (message, variant) => {
    if (!statusElement) {
      return;
    }

    statusElement.textContent = message;
    statusElement.dataset.variant = variant;
  };

  const setSubmitting = (isSubmitting) => {
    if (submitButton instanceof HTMLButtonElement) {
      submitButton.disabled = isSubmitting;
      submitButton.textContent = isSubmitting ? "Sending..." : "Send Message";
    }
    form.setAttribute("aria-busy", String(isSubmitting));
  };

  form.addEventListener("submit", async (event) => {
    event.preventDefault();

    const formData = new FormData(form);
    const payload = Object.fromEntries(formData.entries());

    setSubmitting(true);
    setStatus("Sending your message...", "pending");

    try {
      const response = await fetch(submitUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
      });

      const data = await response.json().catch(() => ({}));

      if (!response.ok) {
        const errorMessage = typeof data.error === "string" ? data.error : "Unable to send your message.";
        throw new Error(errorMessage);
      }

      form.reset();
      setStatus("Message sent successfully.", "success");
    } catch (error) {
      setStatus(error instanceof Error ? error.message : "Unable to send your message.", "error");
    } finally {
      setSubmitting(false);
    }
  });
};

window.CORE_READY?.then(() => {
  initContactForm();
});
