const menuToggle = document.getElementById("menuToggle");
const nav = document.getElementById("nav");
const imageMap = window.IMAGE_MAP || {};

document.querySelectorAll("img[data-image-key]").forEach((image) => {
  const key = image.getAttribute("data-image-key");
  if (!key) {
    return;
  }
  const mappedSrc = imageMap[key];
  if (mappedSrc) {
    image.setAttribute("src", mappedSrc);
  }
});

if (menuToggle && nav) {
  const closeNav = () => {
    nav.classList.remove("open");
    menuToggle.setAttribute("aria-expanded", "false");
  };

  const setNavOpen = (isOpen) => {
    nav.classList.toggle("open", isOpen);
    menuToggle.setAttribute("aria-expanded", String(isOpen));
  };

  menuToggle.addEventListener("click", () => {
    const isOpen = !nav.classList.contains("open");
    setNavOpen(isOpen);
  });

  document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
    anchor.addEventListener("click", () => {
      closeNav();
    });
  });

  document.addEventListener("click", (event) => {
    if (!nav.classList.contains("open")) {
      return;
    }
    const target = event.target;
    if (target instanceof Element && !nav.contains(target) && target !== menuToggle) {
      closeNav();
    }
  });

  document.addEventListener("keydown", (event) => {
    if (event.key === "Escape") {
      closeNav();
    }
  });

  window.addEventListener("resize", () => {
    if (window.innerWidth > 860) {
      closeNav();
    }
  });
}
