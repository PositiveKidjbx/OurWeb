const initScrollImageEffects = () => {
  const parallaxPanels = [...document.querySelectorAll(".story-panel")];
  if (!parallaxPanels.length || window.matchMedia("(prefers-reduced-motion: reduce)").matches) {
    return;
  }

  const panelStates = parallaxPanels.map((panel) => ({
    panel,
    current: 0,
    target: 0
  }));

  let isMeasuring = false;
  let isAnimating = false;

  const clamp = (value, min, max) => Math.min(Math.max(value, min), max);
  const easeProgress = (value) => value * value * (3 - (2 * value));

  const measureTargets = () => {
    const viewportHeight = window.innerHeight || 1;

    panelStates.forEach((state) => {
      const { panel } = state;
      const rect = panel.getBoundingClientRect();
      const rawProgress = (viewportHeight - rect.top) / (viewportHeight + rect.height);
      state.target = easeProgress(clamp(rawProgress, 0, 1));
    });

    isMeasuring = false;
  };

  const animateZoom = () => {
    let shouldContinue = false;

    panelStates.forEach((state) => {
      state.current += (state.target - state.current) * 0.12;

      if (Math.abs(state.target - state.current) > 0.001) {
        shouldContinue = true;
      } else {
        state.current = state.target;
      }

      const scale = 1.08;
      const shift = (state.current - 0.5) * -84;

      state.panel.style.setProperty("--image-scale", scale.toFixed(3));
      state.panel.style.setProperty("--image-shift", `${shift.toFixed(1)}px`);
    });

    if (shouldContinue) {
      window.requestAnimationFrame(animateZoom);
      return;
    }

    isAnimating = false;
  };

  const requestUpdate = () => {
    if (!isMeasuring) {
      isMeasuring = true;
      window.requestAnimationFrame(measureTargets);
    }

    if (!isAnimating) {
      isAnimating = true;
      window.requestAnimationFrame(animateZoom);
    }
  };

  measureTargets();
  panelStates.forEach((state) => {
    state.current = state.target;
  });
  animateZoom();
  window.addEventListener("scroll", requestUpdate, { passive: true });
  window.addEventListener("resize", requestUpdate);
};

const initSustainabilityCarousel = () => {
  const carousel = document.querySelector("[data-sustainability-carousel]");
  if (!carousel) {
    return;
  }

  const track = carousel.querySelector("[data-carousel-track]");
  const previousButton = carousel.querySelector("[data-carousel-prev]");
  const nextButton = carousel.querySelector("[data-carousel-next]");
  const cards = [...carousel.querySelectorAll(".sustainability-card")];

  if (!track || !previousButton || !nextButton || !cards.length) {
    return;
  }

  let activeIndex = 0;

  const getVisibleCount = () => {
    if (window.innerWidth <= 640) {
      return 1;
    }

    if (window.innerWidth <= 980) {
      return 2;
    }

    return 3;
  };

  const updateCarousel = () => {
    const visibleCount = getVisibleCount();
    const maxIndex = Math.max(cards.length - visibleCount, 0);
    activeIndex = Math.min(Math.max(activeIndex, 0), maxIndex);

    const cardWidth = cards[0].getBoundingClientRect().width;
    const gap = parseFloat(window.getComputedStyle(track).columnGap) || 0;
    const offset = activeIndex * (cardWidth + gap);

    track.style.transform = `translateX(-${offset}px)`;
    previousButton.disabled = activeIndex === 0;
    nextButton.disabled = activeIndex === maxIndex;
  };

  previousButton.addEventListener("click", () => {
    activeIndex -= 1;
    updateCarousel();
  });

  nextButton.addEventListener("click", () => {
    activeIndex += 1;
    updateCarousel();
  });

  window.addEventListener("resize", updateCarousel);
  updateCarousel();
};

window.CORE_READY?.then(() => {
  initScrollImageEffects();
  initSustainabilityCarousel();
});
