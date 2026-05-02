const getRootPath = () => document.body.dataset.root || "./";

const withRoot = (value) => value.replaceAll("{{root}}", getRootPath());

const loadIncludes = async () => {
  const includeTargets = [...document.querySelectorAll("[data-include]")];

  await Promise.all(includeTargets.map(async (target) => {
    const includePath = target.getAttribute("data-include");
    if (!includePath) {
      return;
    }

    const response = await fetch(withRoot(includePath));
    if (!response.ok) {
      throw new Error(`Failed to load include: ${includePath}`);
    }

    target.outerHTML = withRoot(await response.text());
  }));
};

const applyImageMap = () => {
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
};

const initNavigation = () => {
  const menuToggle = document.getElementById("menuToggle");
  const nav = document.getElementById("nav");

  if (!menuToggle || !nav) {
    return;
  }

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

  document.querySelectorAll('a[href^="#"], a[href$=".html"], a[href$="/"]').forEach((anchor) => {
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

  const currentPath = window.location.pathname.split("/").pop() || "index.html";
  nav.querySelectorAll("a[href]").forEach((anchor) => {
    const href = anchor.getAttribute("href") || "";
    const normalizedHref = href.replaceAll("{{root}}", "").split("/").pop() || "index.html";
    if (normalizedHref === currentPath) {
      anchor.setAttribute("aria-current", "page");
    }
  });
};

const initRevealAnimations = () => {
  const revealTargets = [...document.querySelectorAll(".reveal, .reveal-title")];
  if (!revealTargets.length) {
    return;
  }

  if (window.matchMedia("(prefers-reduced-motion: reduce)").matches || !("IntersectionObserver" in window)) {
    revealTargets.forEach((target) => target.classList.add("is-visible"));
    return;
  }

  const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (!entry.isIntersecting) {
        return;
      }

      entry.target.classList.add("is-visible");
      observer.unobserve(entry.target);
    });
  }, {
    threshold: 0.45
  });

  revealTargets.forEach((target) => observer.observe(target));
};

const initPage = async () => {
  try {
    await loadIncludes();
  } catch (error) {
    console.error(error);
  }

  applyImageMap();
  initNavigation();
  initRevealAnimations();
};

window.CORE_READY = initPage();
