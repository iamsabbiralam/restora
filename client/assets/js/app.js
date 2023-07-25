var options = {
  series: [
    {
      name: "Orders",
      data: [23, 55, 22, 45, 20, 32, 22, 42, 21, 44, 22, 30],
    },
    {
      name: "Sales",
      data: [40, 35, 66, 28, 38, 55, 45, 70, 55, 69, 46, 49],
    },
  ],
  chart: {
    animations: {
      enabled: false,
    },
    height: 310,
    type: "area",
    zoom: {
      enabled: false,
    },
    toolbar: {
      show: false,
    },
  },
  dataLabels: {
    enabled: false,
  },
  legend: {
    position: "top",
    fontSize: "14px",
    fontWeight: "600",
  },
  stroke: {
    curve: "straight",
    width: "1",
  },
  fill: {
    // type: "solid",
    // opacity: [0.1, 1],
    colors: undefined,
    opacity: [0.1, 1],
    type: ["gradient", "gradient"],
    gradient: {
      shade: "light",
      type: "vertical",
      shadeIntensity: 0.5,
      gradientToColors: undefined,
      inverseColors: true,
      opacityFrom: 0.7,
      opacityTo: 0,
      stops: [0, 80, 100],
      colorStops: [],
    },
  },
  grid: {
    borderColor: "rgba(107 ,114 ,128,0.1)",
  },
  colors: ["rgb(90,102,241)", "rgb(203,213,225)"],
  yaxis: {
    title: {
      style: {
        color: "#adb5be",
        fontSize: "14px",
        fontFamily: "Inter, sans-serif",
        fontWeight: 600,
        cssClass: "apexcharts-yaxis-label",
      },
    },
    labels: {
      style: {
        colors: "rgb(107 ,114 ,128)",
        fontSize: "12px",
      },
      formatter: function (y) {
        return y.toFixed(0) + "";
      },
    },
  },
  xaxis: {
    type: "month",
    categories: [
      "Jan",
      "Feb",
      "Mar",
      "Apr",
      "May",
      "Jun",
      "Jul",
      "Aug",
      "Sep",
      "Oct",
      "Nov",
      "Dec",
    ],
    axisBorder: {
      show: true,
      color: "rgba(119, 119, 142, 0.05)",
      offsetX: 0,
      offsetY: 0,
    },
    axisTicks: {
      show: true,
      borderType: "solid",
      color: "rgba(119, 119, 142, 0.05)",
      width: 6,
      offsetX: 0,
      offsetY: 0,
    },
    labels: {
      rotate: -90,
      style: {
        colors: "rgb(107 ,114 ,128)",
        fontSize: "12px",
      },
    },
  },
};
var chart = new ApexCharts(document.querySelector("#earnings"), options);
chart.render();
function Earnings() {
  chart.updateOptions({
    colors: ["rgb(" + myVarVal + ")", "rgb(203,213,225)"],
  });
}

if (localStorage.Syntodarktheme) {
  document.querySelector("html").setAttribute("class", "dark");
  document.querySelector("html").setAttribute("data-header-styles", "dark");
}
if (localStorage.Syntortl) {
  document.querySelector("html").setAttribute("dir", "rtl");
}
if (localStorage.Syntolayout == "horizontal") {
  document.querySelector("html").setAttribute("data-nav-layout", "horizontal");
}

if (localStorage.Syntolighttheme && localStorage.Syntolayout == "horizontal") {
  document.querySelector("html").setAttribute("data-menu-styles", "light");
}
if (localStorage.Syntoboxed) {
  document.querySelector("html").setAttribute("data-width", "boxed");
}
if (localStorage.Syntoclassic) {
  document.querySelector("html").setAttribute("data-page-style", "classic");
}
if (localStorage.SyntoMenu == "color") {
  document.querySelector("html").setAttribute("data-menu-styles", "color");
}
if (localStorage.SyntoHeader == "color") {
  document.querySelector("html").setAttribute("data-header-styles", "color");
}
if (localStorage.SyntoMenu == "gradient") {
  document.querySelector("html").setAttribute("data-menu-styles", "gradient");
}
if (localStorage.SyntoHeader == "gradient") {
  document.querySelector("html").setAttribute("data-header-styles", "gradient");
}
if (localStorage.SyntoMenu == "dark") {
  document.querySelector("html").setAttribute("data-menu-styles", "dark");
}
if (localStorage.SyntoHeader == "dark") {
  document.querySelector("html").setAttribute("data-header-styles", "dark");
}
if (localStorage.SyntoMenu == "light") {
  document.querySelector("html").setAttribute("data-menu-styles", "light");
}
if (localStorage.SyntoHeader == "light") {
  document.querySelector("html").setAttribute("data-header-styles", "light");
}
if (localStorage.SyntoMenu == "transparent") {
  document
    .querySelector("html")
    .setAttribute("data-menu-styles", "transparent");
}
if (localStorage.SyntoHeader == "transparent") {
  document
    .querySelector("html")
    .setAttribute("data-header-styles", "transparent");
}
if (localStorage.primaryRGB) {
  document
    .querySelector("html")
    .style.setProperty("--color-primary", localStorage.primaryRGB1);
  document
    .querySelector("html")
    .style.setProperty("--color-primary-rgb", localStorage.primaryRGB);
}
if (localStorage.bodyBgRGB) {
  document
    .querySelector("html")
    .style.setProperty("--body-bg", localStorage.bodyBgRGB);
  document
    .querySelector("html")
    .style.setProperty("--dark-bg", localStorage.darkBgRGB);
  let html = document.querySelector("html");
  html.classList.add("dark");
  html.classList.remove("light");
  html.setAttribute("data-menu-styles", "dark");
  html.setAttribute("data-header-styles", "dark");
}
if (localStorage.bgimg == "bgimg1") {
  document.querySelector("html").setAttribute("bg-img", "bgimg1");
}
if (localStorage.bgimg == "bgimg2") {
  document.querySelector("html").setAttribute("bg-img", "bgimg2");
}
if (localStorage.bgimg == "bgimg3") {
  document.querySelector("html").setAttribute("bg-img", "bgimg3");
}
if (localStorage.bgimg == "bgimg4") {
  document.querySelector("html").setAttribute("bg-img", "bgimg4");
}
if (localStorage.bgimg == "bgimg5") {
  document.querySelector("html").setAttribute("bg-img", "bgimg5");
}
if (localStorage.Syntoverticalstyles == "closed") {
  document.querySelector("html").setAttribute("data-vertical-style", "closed");
}
if (localStorage.Syntoverticalstyles == "icontext") {
  document
    .querySelector("html")
    .setAttribute("data-vertical-style", "icontext");
}
if (localStorage.Syntoverticalstyles == "overlay") {
  document.querySelector("html").setAttribute("data-vertical-style", "overlay");
}
if (localStorage.Syntoverticalstyles == "detached") {
  document
    .querySelector("html")
    .setAttribute("data-vertical-style", "detached");
}
if (localStorage.Syntoverticalstyles == "doublemenu") {
  document
    .querySelector("html")
    .setAttribute("data-vertical-style", "doublemenu");
}
/* footer year */
document.getElementById("year").innerHTML = new Date().getFullYear();
/* footer year */

/* back to top */
const scrollToTop = document.querySelector(".scrollToTop");
const $rootElement = document.documentElement;
const $body = document.body;
window.onscroll = () => {
  const scrollTop = window.scrollY || window.pageYOffset;
  const clientHt = $rootElement.scrollHeight - $rootElement.clientHeight;
  if (window.scrollY > 100) {
    scrollToTop.style.display = "flex";
  } else {
    scrollToTop.style.display = "none";
  }
};
scrollToTop.onclick = () => {
  window.scrollTo(0, 0);
};
/* back to top */

if (document.querySelector("#hs-overlay-switcher")) {
  //switcher color pickers
  const pickrContainerPrimary = document.querySelector(
    ".pickr-container-primary"
  );
  const themeContainerPrimary = document.querySelector(
    ".theme-container-primary"
  );
  const pickrContainerBackground = document.querySelector(
    ".pickr-container-background"
  );
  const themeContainerBackground = document.querySelector(
    ".theme-container-background"
  );

  /* for theme primary */
  const nanoThemes = [
    [
      "nano",
      {
        defaultRepresentation: "RGB",
        components: {
          preview: true,
          opacity: false,
          hue: true,

          interaction: {
            hex: false,
            rgba: true,
            hsva: false,
            input: true,
            clear: false,
            save: false,
          },
        },
      },
    ],
  ];
  const nanoButtons = [];
  let nanoPickr = null;
  for (const [theme, config] of nanoThemes) {
    const button = document.createElement("button");
    button.innerHTML = theme;
    nanoButtons.push(button);

    button.addEventListener("click", () => {
      const el = document.createElement("p");
      pickrContainerPrimary.appendChild(el);

      /* Delete previous instance */
      if (nanoPickr) {
        nanoPickr.destroyAndRemove();
      }

      /* Apply active class */
      for (const btn of nanoButtons) {
        btn.classList[btn === button ? "add" : "remove"]("active");
      }

      /* Create fresh instance */
      nanoPickr = new Pickr(
        Object.assign(
          {
            el,
            theme,
            default: "#5e76a6",
          },
          config
        )
      );

      /* Set events */
      nanoPickr.on("changestop", (source, instance) => {
        let color = instance.getColor().toRGBA();
        let html = document.querySelector("html");
        html.style.setProperty(
          "--color-primary",
          `${Math.floor(color[0])} ${Math.floor(color[1])} ${Math.floor(
            color[2]
          )}`
        );
        html.style.setProperty(
          "--color-primary-rgb",
          `${Math.floor(color[0])} ,${Math.floor(color[1])}, ${Math.floor(
            color[2]
          )}`
        );
        /* theme color picker */
        localStorage.setItem(
          "primaryRGB",
          `${Math.floor(color[0])}, ${Math.floor(color[1])}, ${Math.floor(
            color[2]
          )}`
        );
        localStorage.setItem(
          "primaryRGB1",
          `${Math.floor(color[0])} ${Math.floor(color[1])} ${Math.floor(
            color[2]
          )}`
        );
        updateColors();
      });
    });

    themeContainerPrimary.appendChild(button);
  }
  nanoButtons[0].click();
  /* for theme primary */

  /* for theme background */
  const nanoThemes1 = [
    [
      "nano",
      {
        defaultRepresentation: "RGB",
        components: {
          preview: true,
          opacity: false,
          hue: true,

          interaction: {
            hex: false,
            rgba: true,
            hsva: false,
            input: true,
            clear: false,
            save: false,
          },
        },
      },
    ],
  ];
  const nanoButtons1 = [];
  let nanoPickr1 = null;
  for (const [theme, config] of nanoThemes) {
    const button = document.createElement("button");
    button.innerHTML = theme;
    nanoButtons1.push(button);

    button.addEventListener("click", () => {
      const el = document.createElement("p");
      pickrContainerBackground.appendChild(el);

      /* Delete previous instance */
      if (nanoPickr1) {
        nanoPickr1.destroyAndRemove();
      }

      /* Apply active class */
      for (const btn of nanoButtons) {
        btn.classList[btn === button ? "add" : "remove"]("active");
      }

      /* Create fresh instance */
      nanoPickr1 = new Pickr(
        Object.assign(
          {
            el,
            theme,
            default: "#5e76a6",
          },
          config
        )
      );

      /* Set events */
      nanoPickr1.on("changestop", (source, instance) => {
        let color = instance.getColor().toRGBA();
        let html = document.querySelector("html");
        html.style.setProperty(
          "--body-bg",
          `${Math.floor(color[0] + 14)}
           ${Math.floor(color[1] + 14)}
            ${Math.floor(color[2] + 14)}`
        );
        html.style.setProperty(
          "--dark-bg",
          `
          ${Math.floor(color[0])}
          ${Math.floor(color[1])}
          ${Math.floor(color[2])}

          `
        );
        localStorage.removeItem("bgtheme");
        updateColors();
        html.classList.add("dark");
        html.classList.remove("light");
        html.setAttribute("data-menu-styles", "dark");
        html.setAttribute("data-header-styles", "dark");
        document.querySelector("#switcher-dark-theme").checked = true;
        localStorage.setItem(
          "bodyBgRGB",
          `${Math.floor(color[0] + 14)}
           ${Math.floor(color[1] + 14)}
            ${Math.floor(color[2] + 14)}`
        );
        localStorage.setItem(
          "darkBgRGB",
          `${Math.floor(color[0])} ${Math.floor(color[1])} ${Math.floor(
            color[2]
          )}`
        );
      });
    });
    themeContainerBackground.appendChild(button);
  }
  nanoButtons1[0].click();
  /* for theme background */
}

let mainContent;
(function () {
  let html = document.querySelector("html");
  mainContent = document.querySelector(".main-content");
  localStorageBackup();
  if (document.querySelector("#hs-overlay-switcher")) {
    switcherClick();
    checkOptions();
  }
  /* LTR to RTL */
  // html.setAttribute("dir" , "rtl") // for rtl version
})();

function ltrFn() {
  let html = document.querySelector("html");
  html.setAttribute("dir", "ltr");
  document.querySelector("#switcher-ltr").checked = true;
  checkOptions();
}

function rtlFn() {
  let html = document.querySelector("html");
  html.setAttribute("dir", "rtl");
  checkOptions();
}

function darkFn() {
  let html = document.querySelector("html");
  html.classList.add("dark");
  html.classList.remove("light");
  document
    .querySelector("html")
    .style.removeProperty("--body-bg", localStorage.bodyBgRGB);
  document
    .querySelector("html")
    .style.removeProperty("--dark-bg", localStorage.darkBgRGB);
  updateColors();
  localStorage.setItem("Syntodarktheme", true);
  localStorage.removeItem("Syntolighttheme");
  checkOptions();
  html.style.removeProperty("--color-primary");
  html.style.removeProperty("--color-primary-rgb");
}

function checkOptions() {
  // dark
  if (localStorage.getItem("Syntodarktheme")) {
    document.querySelector("#switcher-dark-theme").checked = true;
  }
  // light
  if (localStorage.getItem("Syntolighttheme")) {
    document.querySelector("#switcher-light-theme").checked = true;
  }

  //RTL
  if (localStorage.getItem("Syntortl")) {
    document.querySelector("#switcher-rtl").checked = true;
  }
  if (localStorage.getItem("Syntoltr")) {
    document.querySelector("#switcher-ltr").checked = true;
  }
}

// chart colors
let myVarVal, primaryRGB1;
updateColors();

function localStorageBackup() {
  // if there is a value stored, update color picker and background color
  // Used to retrive the data from local storage
  if (localStorage.primaryRGB) {
    if (document.querySelector(".theme-container-primary")) {
      document.querySelector(".theme-container-primary").value =
        localStorage.primaryRGB;
    }
    document
      .querySelector("html")
      .style.setProperty("--color-primary", localStorage.primaryRGB1);
    document
      .querySelector("html")
      .style.setProperty("--color-primary-rgb", localStorage.primaryRGB);
  }
  if (localStorage.bodyBgRGB) {
    if (document.querySelector(".theme-container-background")) {
      document.querySelector(".theme-container-background").value =
        localStorage.bodyBgRGB;
    }
    document
      .querySelector("html")
      .style.setProperty("--body-bg", localStorage.bodyBgRGB);
    document
      .querySelector("html")
      .style.setProperty("--dark-bg", localStorage.darkBgRGB);
    let html = document.querySelector("html");
    html.classList.add("dark");
    html.classList.remove("light");
    html.setAttribute("data-menu-styles", "dark");
    html.setAttribute("data-header-styles", "dark");
    document.querySelector("#switcher-dark-theme").checked = true;
  }
  if (localStorage.Syntodarktheme) {
    let html = document.querySelector("html");
    html.classList.add("dark");
    html.classList.remove("light");
  }
  if (localStorage.Syntortl) {
    let html = document.querySelector("html");
    setTimeout(() => {
      html.setAttribute("dir", "rtl");
    }, 100);
  }
}

// for menu target scroll on click
window.addEventListener("scroll", reveal);
function reveal() {
  var reveals = document.querySelectorAll(".reveal");

  for (var i = 0; i < reveals.length; i++) {
    var windowHeight = window.innerHeight;
    var cardTop = reveals[i].getBoundingClientRect().top;
    var cardRevealPoint = 150;

    if (cardTop < windowHeight - cardRevealPoint) {
      reveals[i].classList.add("active");
    } else {
      reveals[i].classList.remove("active");
    }
  }
}
reveal();
const pageLink = document.querySelectorAll(".side-menu__item");
pageLink.forEach((elem) => {
  if (elem != "javascript:void(0);" && elem !== "#") {
    elem.addEventListener("click", (e) => {
      e.preventDefault();
      document.querySelector(elem.getAttribute("href"))?.scrollIntoView({
        behavior: "smooth",
        offsetTop: 1 - 60,
      });
    });
  }
});
// section menu active
function onScroll(event) {
  const sections = document.querySelectorAll(".side-menu__item");
  const scrollPos =
    window.pageYOffset ||
    document.documentElement.scrollTop ||
    document.body.scrollTop;

  sections.forEach((elem) => {
    const val = elem.getAttribute("href");
    let refElement;
    if (val != "javascript:void(0);" && val !== "#") {
      refElement = document.querySelector(val);
    }
    const scrollTopMinus = scrollPos + 73;
    if (
      refElement?.offsetTop <= scrollTopMinus &&
      refElement?.offsetTop + refElement.offsetHeight > scrollTopMinus
    ) {
      if (elem.parentElement.parentElement.classList.contains("child1")) {
        elem.parentElement.parentElement.parentElement.children[0].classList.add(
          "active"
        );
      }
      elem.classList.add("active");
      if (elem.closest(".child1")?.previousElementSibling) {
        elem.closest(".child1").previousElementSibling.classList.add("active");
      }
    } else {
      elem.classList.remove("active");
    }
  });
}
window.document.addEventListener("scroll", onScroll);
// for menu target scroll on click

/* count-up */
var i = 1;
setInterval(() => {
  document.querySelectorAll(".count-up").forEach((ele) => {
    if (ele.getAttribute("data-count") >= i) {
      i = i + 1;
      ele.innerText = i;
    }
  });
}, 50);
/* count-up */

/* Swiper 8 */
var swiper8 = new Swiper(".testimonials-swipe", {
  autoplay: {
    delay: 2000,
    disableOnInteraction: false,
  },
  slidesPerView: 1,
  spaceBetween: 30,
  watchSlidesProgress: true,
  freeMode: true,
  breakpoints: {
    768: {
      slidesPerView: 2,
      spaceBetween: 40,
    },
    1024: {
      slidesPerView: 2,
      spaceBetween: 50,
    },
    1400: {
      slidesPerView: 4,
      spaceBetween: 50,
    },
  },
});

let html = document.querySelector("html");
html.setAttribute("dir", "ltr");
html.setAttribute("data-nav-layout", "horizontal");
html.setAttribute("data-nav-style", "menu-click");
html.setAttribute("data-menu-position", "fixed");

("use strict");
(() => {
  var navbar1 = document.querySelector(".app-sidebar");
  var sticky1 = navbar1.clientHeight;
  function stickyFn() {
    if (window.pageYOffset > 2) {
      navbar1.classList.add("sticky-pin");
    } else {
      navbar1.classList.remove("sticky-pin");
    }
  }
  window.addEventListener("scroll", stickyFn);
  window.addEventListener("DOMContentLoaded", stickyFn);
})();

(function () {
  "use strict";

  /* footer year */
  document.getElementById("year").innerHTML = new Date().getFullYear();
  /* footer year */

  /*Start Sidemenu Scroll */
  var myElement = document.getElementById("sidebar-scroll");
  new SimpleBar(myElement, { autoHide: true });
  /*End Sidemenu Scroll */

  if (document.querySelector("#hs-overlay-switcher")) {
    /*Start Switcher Scroll */
    var myElement1 = document.getElementById("switcher-body");
    new SimpleBar(myElement1, { autoHide: true });
    /*End Switcher Scroll */

    //switcher color pickers
    const pickrContainerPrimary = document.querySelector(
      ".pickr-container-primary"
    );
    const themeContainerPrimary = document.querySelector(
      ".theme-container-primary"
    );
    const pickrContainerBackground = document.querySelector(
      ".pickr-container-background"
    );
    const themeContainerBackground = document.querySelector(
      ".theme-container-background"
    );

    /* for theme primary */
    const nanoThemes = [
      [
        "nano",
        {
          defaultRepresentation: "RGB",
          components: {
            preview: true,
            opacity: false,
            hue: true,

            interaction: {
              hex: false,
              rgba: true,
              hsva: false,
              input: true,
              clear: false,
              save: false,
            },
          },
        },
      ],
    ];
    const nanoButtons = [];
    let nanoPickr = null;
    for (const [theme, config] of nanoThemes) {
      const button = document.createElement("button");
      button.innerHTML = theme;
      nanoButtons.push(button);

      button.addEventListener("click", () => {
        const el = document.createElement("p");
        pickrContainerPrimary.appendChild(el);

        /* Delete previous instance */
        if (nanoPickr) {
          nanoPickr.destroyAndRemove();
        }

        /* Apply active class */
        for (const btn of nanoButtons) {
          btn.classList[btn === button ? "add" : "remove"]("active");
        }

        /* Create fresh instance */
        nanoPickr = new Pickr(
          Object.assign(
            {
              el,
              theme,
              default: "#5e76a6",
            },
            config
          )
        );

        /* Set events */
        nanoPickr.on("changestop", (source, instance) => {
          let color = instance.getColor().toRGBA();
          let html = document.querySelector("html");
          html.style.setProperty(
            "--color-primary",
            `${Math.floor(color[0])} ${Math.floor(color[1])} ${Math.floor(
              color[2]
            )}`
          );
          html.style.setProperty(
            "--color-primary-rgb",
            `${Math.floor(color[0])} ,${Math.floor(color[1])}, ${Math.floor(
              color[2]
            )}`
          );
          /* theme color picker */
          localStorage.setItem(
            "primaryRGB",
            `${Math.floor(color[0])}, ${Math.floor(color[1])}, ${Math.floor(
              color[2]
            )}`
          );
          localStorage.setItem(
            "primaryRGB1",
            `${Math.floor(color[0])} ${Math.floor(color[1])} ${Math.floor(
              color[2]
            )}`
          );
          updateColors();
        });
      });

      themeContainerPrimary.appendChild(button);
    }
    nanoButtons[0].click();
    /* for theme primary */

    /* for theme background */
    const nanoThemes1 = [
      [
        "nano",
        {
          defaultRepresentation: "RGB",
          components: {
            preview: true,
            opacity: false,
            hue: true,

            interaction: {
              hex: false,
              rgba: true,
              hsva: false,
              input: true,
              clear: false,
              save: false,
            },
          },
        },
      ],
    ];
    const nanoButtons1 = [];
    let nanoPickr1 = null;
    for (const [theme, config] of nanoThemes) {
      const button = document.createElement("button");
      button.innerHTML = theme;
      nanoButtons1.push(button);

      button.addEventListener("click", () => {
        const el = document.createElement("p");
        pickrContainerBackground.appendChild(el);

        /* Delete previous instance */
        if (nanoPickr1) {
          nanoPickr1.destroyAndRemove();
        }

        /* Apply active class */
        for (const btn of nanoButtons) {
          btn.classList[btn === button ? "add" : "remove"]("active");
        }

        /* Create fresh instance */
        nanoPickr1 = new Pickr(
          Object.assign(
            {
              el,
              theme,
              default: "#5e76a6",
            },
            config
          )
        );

        /* Set events */
        nanoPickr1.on("changestop", (source, instance) => {
          let color = instance.getColor().toRGBA();
          let html = document.querySelector("html");
          html.style.setProperty(
            "--body-bg",
            `${Math.floor(color[0] + 14)}
             ${Math.floor(color[1] + 14)}
              ${Math.floor(color[2] + 14)}`
          );
          html.style.setProperty(
            "--dark-bg",
            `
            ${Math.floor(color[0])}
            ${Math.floor(color[1])}
            ${Math.floor(color[2])}
  
            `
          );
          localStorage.removeItem("bgtheme");
          updateColors();
          html.classList.add("dark");
          html.classList.remove("light");
          html.setAttribute("data-menu-styles", "dark");
          html.setAttribute("data-header-styles", "dark");
          document.querySelector("#switcher-dark-theme").checked = true;
          localStorage.setItem(
            "bodyBgRGB",
            `${Math.floor(color[0] + 14)}
             ${Math.floor(color[1] + 14)}
              ${Math.floor(color[2] + 14)}`
          );
          localStorage.setItem(
            "darkBgRGB",
            `${Math.floor(color[0])} ${Math.floor(color[1])} ${Math.floor(
              color[2]
            )}`
          );
        });
      });
      themeContainerBackground.appendChild(button);
    }
    nanoButtons1[0].click();
    /* for theme background */
  }

  /* box with close button */
  let DIV_Box = ".box";
  let boxRemoveBtn = document.querySelectorAll(".box-remove");
  boxRemoveBtn.forEach((ele) => {
    ele.addEventListener("click", function (e) {
      e.preventDefault();
      let $this = this;
      let box = $this.closest(DIV_Box);
      box.remove();
      return false;
    });
  });
  /* box with close button */

  /* box with fullscreen */
  let boxFullscreenBtn = document.querySelectorAll(".box-fullscreen");
  boxFullscreenBtn.forEach((ele) => {
    ele.addEventListener("click", function (e) {
      let $this = this;
      let box = $this.closest(DIV_Box);
      box.classList.toggle("box-fullscreen");
      box.classList.remove("box-collapsed");
      e.preventDefault();
      return false;
    });
  });
  /* box with fullscreen */ /* back to top */
  const scrollToTop = document.querySelector(".scrollToTop");
  const $rootElement = document.documentElement;
  const $body = document.body;
  window.onscroll = () => {
    const scrollTop = window.scrollY || window.pageYOffset;
    const clientHt = $rootElement.scrollHeight - $rootElement.clientHeight;
    if (window.scrollY > 100) {
      scrollToTop.style.display = "flex";
    } else {
      scrollToTop.style.display = "none";
    }
  };
  scrollToTop.onclick = () => {
    window.scrollTo(0, 0);
  };
  /* back to top */

  /*header-remove */
  const headerbtn = document.querySelectorAll(".header-remove-btn");

  headerbtn.forEach((button, index) => {
    button.addEventListener("click", (e) => {
      e.preventDefault();
      e.stopPropagation();
      button.parentNode.remove();
      if (document.getElementById("allCartsContainer")) {
        document.getElementById("cart-data").innerText = `${
          document.getElementById("allCartsContainer").children.length
        } Items`;
        document.getElementById("cart-data2").innerText = `${
          document.getElementById("allCartsContainer").children.length
        }`;
      }
      if (document.getElementById("allNotifyContainer")) {
        document.getElementById("notify-data").innerText = `${
          document.getElementById("allNotifyContainer").children.length
        }`;
      }

      if (document.getElementById("allCartsContainer")) {
        if (document.getElementById("allCartsContainer").children.length == 0) {
          document.getElementById("allCartsContainer").parentNode.innerHTML = `
                        <div class="p-6 pb-8 text-center">
                          <div>
                            <i class="ri ri-shopping-cart-2-line leading-none text-4xl avatar avatar-lg bg-primary/20 text-primary rounded-full p-3 align-middle flex justify-center mx-auto"></i>
                            <div class="mt-5">
                              <p class="text-base font-semibold text-gray-800 dark:text-white mb-1">
                                No Items In Cart
                              </p>
                              <p class="text-xs text-gray-500 dark:text-white/70">
                              When you have Items added here , they will appear here.
                              </p>
                              <a href="javascript:void(0);" class="m-0 ti-btn ti-btn-primary py-1 mt-5"><i class="ti ti-arrow-right text-base leading-none"></i>Continue Shopping</a>
                            </div>
                          </div>
                        </div>`;
        }
      }
      if (document.getElementById("allNotifyContainer")) {
        if (
          document.getElementById("allNotifyContainer").children.length == 0
        ) {
          document.getElementById("allNotifyContainer").parentNode.innerHTML = `
          <div class="p-6 pb-8 text-center">
          <div>
            <i class="ri ri-notification-off-line leading-none text-4xl avatar avatar-lg bg-secondary/20 text-secondary rounded-full p-3 align-middle flex justify-center mx-auto"></i>
            <div class="mt-5">
              <p class="text-base font-semibold text-gray-800 dark:text-white mb-1">
                No Notifications Found
              </p>
              <p class="text-xs text-gray-500 dark:text-white/70">
              When you have notifications added here , they will appear here.
              </p>
            </div>
          </div>
        </div>`;
        }
      }
    });
  });
  /*header-remove */
})();

/* full screen */
var elem = document.documentElement;
function openFullscreen() {
  let open = document.querySelector(".full-screen-open");
  let close = document.querySelector(".full-screen-close");
  if (!document.fullscreenElement) {
    if (elem.requestFullscreen) {
      elem.requestFullscreen();
    } else if (elem.webkitRequestFullscreen) {
      /* Safari */
      elem.webkitRequestFullscreen();
    } else if (elem.msRequestFullscreen) {
      /* IE11 */
      elem.msRequestFullscreen();
    }
    close.classList.add("block");
    close.classList.remove("hidden");
    open.classList.add("hidden");
  } else {
    if (document.exitFullscreen) {
      document.exitFullscreen();
    } else if (document.webkitExitFullscreen) {
      /* Safari */
      document.webkitExitFullscreen();
    } else if (document.msExitFullscreen) {
      /* IE11 */
      document.msExitFullscreen();
    }
    close.classList.remove("block");
    open.classList.remove("hidden");
    close.classList.add("hidden");
    open.classList.add("block");
  }
}

("use strict");

(function () {
  let html = document.querySelector("html");
  mainContent = document.querySelector(".main-content");

  localStorageBackup();
  if (document.querySelector("#hs-overlay-switcher")) {
    switcherClick();
    checkOptions();
    setTimeout(() => {
      checkOptions();
    }, 1000);
  }
})();

function switcherClick() {
  let ltrBtn,
    rtlBtn,
    verticalBtn,
    horiBtn,
    lightBtn,
    darkBtn,
    boxedBtn,
    fullwidthBtn,
    scrollHeaderBtn,
    scrollMenuBtn,
    fixedHeaderBtn,
    fixedMenuBtn,
    lightHeaderBtn,
    darkHeaderBtn,
    colorHeaderBtn,
    gradientHeaderBtn,
    lightMenuBtn,
    darkMenuBtn,
    colorMenuBtn,
    gradientMenuBtn,
    transparentMenuBtn,
    transparentHeaderBtn,
    regular,
    classic,
    defaultBtn,
    closedBtn,
    iconTextBtn,
    detachedBtn,
    overlayBtn,
    doubleBtn,
    resetBtn,
    menuClickBtn,
    menuHoverBtn,
    iconClickBtn,
    iconHoverBtn,
    primaryDefaultColor1Btn,
    primaryDefaultColor2Btn,
    primaryDefaultColor3Btn,
    primaryDefaultColor4Btn,
    primaryDefaultColor5Btn,
    bgDefaultColor1Btn,
    bgDefaultColor2Btn,
    bgDefaultColor3Btn,
    bgDefaultColor4Btn,
    bgDefaultColor5Btn,
    bgImage1Btn,
    bgImage2Btn,
    bgImage3Btn,
    bgImage4Btn,
    bgImage5Btn,
    ResetAll;
  let html = document.querySelector("html");
  lightBtn = document.querySelector("#switcher-light-theme");
  darkBtn = document.querySelector("#switcher-dark-theme");
  ltrBtn = document.querySelector("#switcher-ltr");
  rtlBtn = document.querySelector("#switcher-rtl");
  verticalBtn = document.querySelector("#switcher-vertical");
  horiBtn = document.querySelector("#switcher-horizontal");
  boxedBtn = document.querySelector("#switcher-boxed");
  fullwidthBtn = document.querySelector("#switcher-full-width");
  fixedMenuBtn = document.querySelector("#switcher-menu-fixed");
  scrollMenuBtn = document.querySelector("#switcher-menu-scroll");
  fixedHeaderBtn = document.querySelector("#switcher-header-fixed");
  scrollHeaderBtn = document.querySelector("#switcher-header-scroll");
  lightHeaderBtn = document.querySelector("#switcher-header-light");
  darkHeaderBtn = document.querySelector("#switcher-header-dark");
  colorHeaderBtn = document.querySelector("#switcher-header-primary");
  gradientHeaderBtn = document.querySelector("#switcher-header-gradient");
  transparentHeaderBtn = document.querySelector("#switcher-header-transparent");
  lightMenuBtn = document.querySelector("#switcher-menu-light");
  darkMenuBtn = document.querySelector("#switcher-menu-dark");
  colorMenuBtn = document.querySelector("#switcher-menu-primary");
  gradientMenuBtn = document.querySelector("#switcher-menu-gradient");
  transparentMenuBtn = document.querySelector("#switcher-menu-transparent");
  regular = document.querySelector("#switcher-regular");
  classic = document.querySelector("#switcher-classic");
  defaultBtn = document.querySelector("#switcher-default-menu");
  menuClickBtn = document.querySelector("#switcher-menu-click");
  menuHoverBtn = document.querySelector("#switcher-menu-hover");
  iconClickBtn = document.querySelector("#switcher-icon-click");
  iconHoverBtn = document.querySelector("#switcher-icon-hover");
  closedBtn = document.querySelector("#switcher-closed-menu");
  iconTextBtn = document.querySelector("#switcher-icontext-menu");
  overlayBtn = document.querySelector("#switcher-icon-overlay");
  doubleBtn = document.querySelector("#switcher-double-menu");
  detachedBtn = document.querySelector("#switcher-detached");
  resetBtn = document.querySelector("#resetbtn");
  primaryDefaultColor1Btn = document.querySelector("#switcher-primary");
  primaryDefaultColor2Btn = document.querySelector("#switcher-primary1");
  primaryDefaultColor3Btn = document.querySelector("#switcher-primary2");
  primaryDefaultColor4Btn = document.querySelector("#switcher-primary3");
  primaryDefaultColor5Btn = document.querySelector("#switcher-primary4");
  bgDefaultColor1Btn = document.querySelector("#switcher-background");
  bgDefaultColor2Btn = document.querySelector("#switcher-background1");
  bgDefaultColor3Btn = document.querySelector("#switcher-background2");
  bgDefaultColor4Btn = document.querySelector("#switcher-background3");
  bgDefaultColor5Btn = document.querySelector("#switcher-background4");
  bgImage1Btn = document.querySelector("#switcher-bg-img");
  bgImage2Btn = document.querySelector("#switcher-bg-img1");
  bgImage3Btn = document.querySelector("#switcher-bg-img2");
  bgImage4Btn = document.querySelector("#switcher-bg-img3");
  bgImage5Btn = document.querySelector("#switcher-bg-img4");
  ResetAll = document.querySelector("#reset-all");

  // primary theme
  let primaryColor1Var = primaryDefaultColor1Btn.addEventListener(
    "click",
    () => {
      localStorage.setItem("primaryRGB", "58, 88, 146");
      localStorage.setItem("primaryRGB1", "58 88 146");
      html.style.setProperty("--color-primary-rgb", `58, 88, 146`);
      html.style.setProperty("--color-primary", `58 88 146`);
      updateColors();
    }
  );
  let primaryColor2Var = primaryDefaultColor2Btn.addEventListener(
    "click",
    () => {
      localStorage.setItem("primaryRGB", "92, 144, 163");
      localStorage.setItem("primaryRGB1", "92 144 163");
      html.style.setProperty("--color-primary-rgb", `92, 144, 163`);
      html.style.setProperty("--color-primary", `92 144 163`);
      updateColors();
    }
  );
  let primaryColor3Var = primaryDefaultColor3Btn.addEventListener(
    "click",
    () => {
      localStorage.setItem("primaryRGB", "172, 172, 80");
      localStorage.setItem("primaryRGB1", "172 172 80");
      html.style.setProperty("--color-primary-rgb", `172, 172, 80`);
      html.style.setProperty("--color-primary", `172 172 80`);
      updateColors();
    }
  );
  let primaryColor4Var = primaryDefaultColor4Btn.addEventListener(
    "click",
    () => {
      localStorage.setItem("primaryRGB", "165, 94, 131");
      localStorage.setItem("primaryRGB1", "165 94 131");
      html.style.setProperty("--color-primary-rgb", `165, 94, 131`);
      html.style.setProperty("--color-primary", `165 94 131`);
      updateColors();
    }
  );
  let primaryColor5Var = primaryDefaultColor5Btn.addEventListener(
    "click",
    () => {
      localStorage.setItem("primaryRGB", "87, 68, 117");
      localStorage.setItem("primaryRGB1", "87 68 117");
      html.style.setProperty("--color-primary-rgb", `87, 68, 117`);
      html.style.setProperty("--color-primary", `87 68 117`);
      updateColors();
    }
  );

  // Background theme
  let backgroundColor1Var = bgDefaultColor1Btn.addEventListener("click", () => {
    localStorage.setItem("bodyBgRGB", `${50 + 14} ${62 + 14} ${93 + 14}`);
    localStorage.setItem("darkBgRGB", "50 62 93");
    localStorage.removeItem("SyntoHeader");
    localStorage.removeItem("SyntoMenu");
    html.classList.add("dark");
    html.classList.remove("light");
    html.setAttribute("data-menu-styles", "dark");
    html.setAttribute("data-header-styles", "dark");
    document
      .querySelector("html")
      .style.setProperty("--body-bg", localStorage.bodyBgRGB);
    document
      .querySelector("html")
      .style.setProperty("--dark-bg", localStorage.darkBgRGB);
    document.querySelector("#switcher-dark-theme").checked = true;
  });
  let backgroundColor2Var = bgDefaultColor2Btn.addEventListener("click", () => {
    localStorage.setItem("bodyBgRGB", `${81 + 14} ${93 + 14} ${50 + 14}`);
    localStorage.setItem("darkBgRGB", "81 93 50");
    localStorage.removeItem("SyntoHeader");
    localStorage.removeItem("SyntoMenu");
    html.classList.add("dark");
    html.classList.remove("light");
    html.setAttribute("data-menu-styles", "dark");
    html.setAttribute("data-header-styles", "dark");
    document
      .querySelector("html")
      .style.setProperty("--body-bg", localStorage.bodyBgRGB);
    document
      .querySelector("html")
      .style.setProperty("--dark-bg", localStorage.darkBgRGB);
    document.querySelector("#switcher-dark-theme").checked = true;
  });
  let backgroundColor3Var = bgDefaultColor3Btn.addEventListener("click", () => {
    localStorage.setItem("bodyBgRGB", `${79 + 14} ${50 + 14} ${93 + 14}`);
    localStorage.setItem("darkBgRGB", "79 50 93");
    localStorage.removeItem("SyntoHeader");
    localStorage.removeItem("SyntoMenu");
    html.classList.add("dark");
    html.classList.remove("light");
    html.setAttribute("data-menu-styles", "dark");
    html.setAttribute("data-header-styles", "dark");
    document
      .querySelector("html")
      .style.setProperty("--body-bg", localStorage.bodyBgRGB);
    document
      .querySelector("html")
      .style.setProperty("--dark-bg", localStorage.darkBgRGB);
    document.querySelector("#switcher-dark-theme").checked = true;
  });
  let backgroundColor4Var = bgDefaultColor4Btn.addEventListener("click", () => {
    localStorage.setItem("bodyBgRGB", `${50 + 14} ${87 + 14} ${93 + 14}`);
    localStorage.setItem("darkBgRGB", "50 87 93");
    localStorage.removeItem("SyntoHeader");
    localStorage.removeItem("SyntoMenu");
    html.classList.add("dark");
    html.classList.remove("light");
    html.setAttribute("data-menu-styles", "dark");
    html.setAttribute("data-header-styles", "dark");
    document
      .querySelector("html")
      .style.setProperty("--body-bg", localStorage.bodyBgRGB);
    document
      .querySelector("html")
      .style.setProperty("--dark-bg", localStorage.darkBgRGB);
    document.querySelector("#switcher-dark-theme").checked = true;
  });
  let backgroundColor5Var = bgDefaultColor5Btn.addEventListener("click", () => {
    localStorage.setItem("bodyBgRGB", `${93 + 14} ${50 + 14} ${50 + 14}`);
    localStorage.setItem("darkBgRGB", "93 50 50");
    localStorage.removeItem("SyntoHeader");
    localStorage.removeItem("SyntoMenu");
    html.classList.add("dark");
    html.classList.remove("light");
    html.setAttribute("data-menu-styles", "dark");
    html.setAttribute("data-header-styles", "dark");
    document
      .querySelector("html")
      .style.setProperty("--body-bg", localStorage.bodyBgRGB);
    document
      .querySelector("html")
      .style.setProperty("--dark-bg", localStorage.darkBgRGB);
    document.querySelector("#switcher-dark-theme").checked = true;
  });

  // Bg image
  let bgImg1Var = bgImage1Btn.addEventListener("click", () => {
    html.setAttribute("bg-img", "bgimg1");
    localStorage.setItem("bgimg", "bgimg1");
  });
  let bgImg2Var = bgImage2Btn.addEventListener("click", () => {
    html.setAttribute("bg-img", "bgimg2");
    localStorage.setItem("bgimg", "bgimg2");
  });
  let bgImg3Var = bgImage3Btn.addEventListener("click", () => {
    html.setAttribute("bg-img", "bgimg3");
    localStorage.setItem("bgimg", "bgimg3");
  });
  let bgImg4Var = bgImage4Btn.addEventListener("click", () => {
    html.setAttribute("bg-img", "bgimg4");
    localStorage.setItem("bgimg", "bgimg4");
  });
  let bgImg5Var = bgImage5Btn.addEventListener("click", () => {
    html.setAttribute("bg-img", "bgimg5");
    localStorage.setItem("bgimg", "bgimg5");
  });

  /* Light Layout Start */
  let lightThemeVar = lightBtn.addEventListener("click", () => {
    lightFn();
    localStorage.setItem("SyntoHeader", "light");
    // localStorage.setItem("SyntoMenu", "light");
  });
  /* Light Layout End */

  /* Dark Layout Start */
  let darkThemeVar = darkBtn.addEventListener("click", () => {
    darkFn();
    localStorage.setItem("SyntoMenu", "dark");
    localStorage.setItem("SyntoHeader", "dark");
  });
  /* Dark Layout End */

  /* Light Menu Start */
  let lightMenuVar = lightMenuBtn.addEventListener("click", () => {
    html.setAttribute("data-menu-styles", "light");
    localStorage.setItem("SyntoMenu", "light");
  });
  /* Light Menu End */

  /* Color Menu Start */
  let colorMenuVar = colorMenuBtn.addEventListener("click", () => {
    html.setAttribute("data-menu-styles", "color");
    localStorage.setItem("SyntoMenu", "color");
  });
  /* Color Menu End */

  /* Dark Menu Start */
  let darkMenuVar = darkMenuBtn.addEventListener("click", () => {
    html.setAttribute("data-menu-styles", "dark");
    localStorage.setItem("SyntoMenu", "dark");
  });
  /* Dark Menu End */

  /* Gradient Menu Start */
  let gradientMenuVar = gradientMenuBtn.addEventListener("click", () => {
    html.setAttribute("data-menu-styles", "gradient");
    localStorage.setItem("SyntoMenu", "gradient");
  });
  /* Gradient Menu End */

  /* Transparent Menu Start */
  let transparentMenuVar = transparentMenuBtn.addEventListener("click", () => {
    html.setAttribute("data-menu-styles", "transparent");
    localStorage.setItem("SyntoMenu", "transparent");
  });
  /* Transparent Menu End */

  /* Light Header Start */
  let lightHeaderVar = lightHeaderBtn.addEventListener("click", () => {
    html.setAttribute("data-header-styles", "light");
    localStorage.setItem("SyntoHeader", "light");
  });
  /* Light Header End */

  /* Color Header Start */
  let colorHeaderVar = colorHeaderBtn.addEventListener("click", () => {
    html.setAttribute("data-header-styles", "color");
    localStorage.setItem("SyntoHeader", "color");
  });
  /* Color Header End */

  /* Dark Header Start */
  let darkHeaderVar = darkHeaderBtn.addEventListener("click", () => {
    html.setAttribute("data-header-styles", "dark");
    localStorage.setItem("SyntoHeader", "dark");
  });
  /* Dark Header End */

  /* Gradient Header Start */
  let gradientHeaderVar = gradientHeaderBtn.addEventListener("click", () => {
    html.setAttribute("data-header-styles", "gradient");
    localStorage.setItem("SyntoHeader", "gradient");
  });
  /* Gradient Header End */

  /* Transparent Header Start */
  let transparentHeaderVar = transparentHeaderBtn.addEventListener(
    "click",
    () => {
      html.setAttribute("data-header-styles", "transparent");
      localStorage.setItem("SyntoHeader", "transparent");
    }
  );
  /* Transparent Header End */

  /* Full Width Layout Start */
  let fullwidthVar = fullwidthBtn.addEventListener("click", () => {
    html.setAttribute("data-width", "fullwidth");
    localStorage.setItem("Syntofullwidth", true);
    localStorage.removeItem("Syntoboxed");
  });
  /* Full Width Layout End */

  /* Boxed Layout Start */
  let boxedVar = boxedBtn.addEventListener("click", () => {
    html.setAttribute("data-width", "boxed");
    localStorage.setItem("Syntoboxed", true);
    localStorage.removeItem("Syntofullwidth");
    checkHoriMenu();
  });
  /* Boxed Layout End */

  /* Regular page style Start */
  let shadowVar = regular.addEventListener("click", () => {
    html.setAttribute("data-page-style", "regular");
    localStorage.setItem("Syntoregular", true);
    localStorage.removeItem("Syntoclassic");
  });
  /* Regular page style End */

  /* Classic page style Start */
  let noShadowVar = classic.addEventListener("click", () => {
    html.setAttribute("data-page-style", "classic");
    localStorage.setItem("Syntoclassic", true);
    localStorage.removeItem("Syntoregular");
  });
  /* Classic page style End */

  /* Header-Position Styles Start */
  let fixedHeaderVar = fixedHeaderBtn.addEventListener("click", () => {
    html.setAttribute("data-header-position", "fixed");
    localStorage.setItem("Syntoheaderfixed", true);
    localStorage.removeItem("Syntoheaderscrollable");
  });

  let scrollHeaderVar = scrollHeaderBtn.addEventListener("click", () => {
    html.setAttribute("data-header-position", "scrollable");
    localStorage.setItem("Syntoheaderscrollable", true);
    localStorage.removeItem("Syntoheaderfixed");
  });
  /* Header-Position Styles End */

  /* Menu-Position Styles Start */
  let fixedMenuVar = fixedMenuBtn.addEventListener("click", () => {
    html.setAttribute("data-menu-position", "fixed");
    localStorage.setItem("Syntomenufixed", true);
    localStorage.removeItem("Syntomenuscrollable");
  });

  let scrollMenuVar = scrollMenuBtn.addEventListener("click", () => {
    html.setAttribute("data-menu-position", "scrollable");
    localStorage.setItem("Syntomenuscrollable", true);
    localStorage.removeItem("Syntomenufixed");
  });
  /* Menu-Position Styles End */

  /* Default Sidemenu Start */
  let defaultVar = defaultBtn.addEventListener("click", () => {
    html.setAttribute("data-vertical-style", "default");
    html.setAttribute("data-nav-layout", "vertical");
    toggleSidemenu();
    localStorage.removeItem("Syntoverticalstyles");
  });
  /* Default Sidemenu End */

  /* Closed Sidemenu Start */
  let closedVar = closedBtn.addEventListener("click", () => {
    closedSidemenuFn();
    localStorage.setItem("Syntoverticalstyles", "closed");
  });
  /* Closed Sidemenu End */

  /* Hover Submenu Start */
  let detachedVar = detachedBtn.addEventListener("click", () => {
    detachedFn();
    localStorage.setItem("Syntoverticalstyles", "detached");
  });
  /* Hover Submenu End */

  /* Icon Text Sidemenu Start */
  let iconTextVar = iconTextBtn.addEventListener("click", () => {
    iconTextFn();
    localStorage.setItem("Syntoverticalstyles", "icontext");
  });
  /* Icon Text Sidemenu End */

  /* Icon Overlay Sidemenu Start */
  let overlayVar = overlayBtn.addEventListener("click", () => {
    iconOverayFn();
    localStorage.setItem("Syntoverticalstyles", "overlay");
  });
  /* Icon Overlay Sidemenu End */

  /* doublemenu Sidemenu Start */
  let doubleVar = doubleBtn.addEventListener("click", () => {
    doubletFn();
    localStorage.setItem("Syntoverticalstyles", "doublemenu");
  });
  /* doublemenu Sidemenu End */

  /* Menu Click Sidemenu Start */
  let menuClickVar = menuClickBtn.addEventListener("click", () => {
    html.removeAttribute("data-vertical-style");
    menuClickFn();
    localStorage.setItem("Syntonavstyles", "menu-click");
    localStorage.removeItem("Syntoverticalstyles");
  });
  /* Menu Click Sidemenu End */

  /* Menu Hover Sidemenu Start */
  let menuhoverVar = menuHoverBtn.addEventListener("click", () => {
    html.removeAttribute("data-vertical-style");
    menuhoverFn();
    localStorage.setItem("Syntonavstyles", "menu-hover");
    localStorage.removeItem("Syntoverticalstyles");
  });
  /* Menu Hover Sidemenu End */

  /* icon Click Sidemenu Start */
  let iconClickVar = iconClickBtn.addEventListener("click", () => {
    html.removeAttribute("data-vertical-style");
    iconClickFn();
    localStorage.setItem("Syntonavstyles", "icon-click");
    localStorage.removeItem("Syntoverticalstyles");

    document.querySelector(".main-menu").style.marginLeft = "0px";
    document.querySelector(".main-menu").style.marginRight = "0px";
  });
  /* icon Click Sidemenu End */

  /* icon hover Sidemenu Start */
  let iconhoverVar = iconHoverBtn.addEventListener("click", () => {
    html.removeAttribute("data-vertical-style");
    iconHoverFn();
    localStorage.setItem("Syntonavstyles", "icon-hover");
    localStorage.removeItem("Syntoverticalstyles");

    document.querySelector(".main-menu").style.marginLeft = "0px";
    document.querySelector(".main-menu").style.marginRight = "0px";
  });
  /* icon hover Sidemenu End */

  /* Sidemenu start*/
  let verticalVar = verticalBtn.addEventListener("click", () => {
    let mainContent = document.querySelector(".main-content");
    // local storage
    localStorage.setItem("Syntolayout", "vertical");
    // localStorage.removeItem("Syntolayout");
    // localStorage.setItem("Syntoverticalstyles", 'default');
    verticalFn();
    setNavActive();
    mainContent.removeEventListener("click", clearNavDropdown);
    document.querySelectorAll(".slide").forEach((element) => {
      if (
        element.classList.contains("open") &&
        !element.classList.contains("active")
      ) {
        element.querySelector("ul").style.display = "none";
      }
    });
  });
  /* Sidemenu end */

  /* horizontal start*/
  let horiVar = horiBtn.addEventListener("click", () => {
    let mainContent = document.querySelector(".main-content");
    html.removeAttribute("data-vertical-style");
    //    local storage
    localStorage.setItem("Syntolayout", "horizontal");
    localStorage.removeItem("Syntoverticalstyles");

    horizontalClickFn();
    clearNavDropdown();
    mainContent.addEventListener("click", clearNavDropdown);
  });
  /* horizontal end*/

  /* rtl start */
  let rtlVar = rtlBtn.addEventListener("click", () => {
    localStorage.setItem("Syntortl", true);
    localStorage.removeItem("Syntoltr");
    rtlFn();
    if (document.querySelector("#color-slider")) {
      document.querySelectorAll(".noUi-origin").forEach((e) => {
        e.classList.add("!transform-none");
      });
    }
  });
  /* rtl end */

  /* ltr start */
  let ltrVar = ltrBtn.addEventListener("click", () => {
    //    local storage
    localStorage.setItem("Syntoltr", true);
    localStorage.removeItem("Syntortl");
    ltrFn();
    if (document.querySelector("#color-slider")) {
      document.querySelectorAll(".noUi-origin").forEach((e) => {
        e.classList.remove("!transform-none");
      });
    }
  });
  /* ltr end */

  // reset all start
  let resetVar = ResetAll.addEventListener("click", () => {
    ResetAllFn();
    setNavActive();
    document.querySelectorAll(".slide").forEach((element) => {
      if (
        element.classList.contains("open") &&
        !element.classList.contains("active")
      ) {
        console.log(element);
        element.querySelector("ul").style.display = "none";
      }
    });
  });
  // reset all start
}

function lightFn() {
  let html = document.querySelector("html");
  html.classList.remove("dark");
  html.classList.add("light");
  html.setAttribute("data-header-styles", "light");
  html.setAttribute("data-menu-styles", "dark");
  document.querySelector("#switcher-light-theme").checked = true;
  document.querySelector("#switcher-menu-dark").checked = true;
  document.querySelector("#switcher-header-light").checked = true;
  updateColors();
  localStorage.setItem("Syntolighttheme", true);
  localStorage.removeItem("Syntodarktheme");
  localStorage.removeItem("SyntobgColor");
  localStorage.removeItem("Syntoheaderbg");
  localStorage.removeItem("Syntobgwhite");
  localStorage.removeItem("Syntomenubg");
  localStorage.removeItem("darkBgRGB");
  localStorage.removeItem("bodyBgRGB");
  localStorage.removeItem("SyntoMenu");
  checkOptions();
  html.style.removeProperty("--body-bg");
  html.style.removeProperty("--dark-bg");
  if (
    document.querySelector("html").getAttribute("data-nav-layout") ==
    "horizontal"
  ) {
    document.querySelector("html").setAttribute("data-menu-styles", "light");
    document.querySelector("#switcher-menu-light").checked = true;
  }

  // localStorage.removeItem("SyntoMenu")
}

function verticalFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-layout", "vertical");
  html.setAttribute("data-vertical-style", "overlay");
  html.removeAttribute("data-nav-style");
  localStorage.removeItem("Syntonavstyles");
  html.removeAttribute("toggled");
  document.querySelector("#switcher-vertical").checked = true;
  document.querySelector("#switcher-default-menu").checked = true;

  document.querySelector("#switcher-menu-click").checked = false;
  document.querySelector("#switcher-menu-hover").checked = false;
  document.querySelector("#switcher-icon-click").checked = false;
  document.querySelector("#switcher-icon-hover").checked = false;
  checkOptions();
  document.querySelector(".main-menu").style.marginLeft = "0px";
  document.querySelector(".main-menu").style.marginRight = "0px";
  html.setAttribute("data-menu-styles", "dark");
}

function horizontalClickFn() {
  if (document.querySelector("#switcher-horizontal")) {
    document.querySelector("#switcher-horizontal").checked = true;
  }
  let html = document.querySelector("html");
  html.setAttribute("data-nav-layout", "horizontal");
  if (!html.getAttribute("data-nav-style")) {
    html.setAttribute("data-nav-style", "menu-click");
  }
  if (!localStorage.SyntoMenu && !localStorage.bodyBgRGB) {
    html.setAttribute("data-menu-styles", "light");
  }
  checkOptions();
}

function ResetAllFn() {
  let html = document.querySelector("html");
  checkOptions();

  // clearing localstorage
  localStorage.clear();

  // reseting to light
  lightFn();

  //rangeslider
  if (document.querySelector("#color-slider")) {
    document.querySelectorAll(".noUi-origin").forEach((e) => {
      e.classList.remove("!transform-none");
    });
  }

  // clearing attibutes
  // removing header, menu, pageStyle & boxed
  html.removeAttribute("data-nav-style");
  html.removeAttribute("data-menu-position");
  html.removeAttribute("data-header-position");
  html.removeAttribute("data-width");
  html.removeAttribute("data-page-style");

  // removing theme styles
  html.removeAttribute("bg-img");

  // clear primary & bg color
  html.style.removeProperty(`--color-primary`);
  html.style.removeProperty(`--color-primary-rgb`);
  html.style.removeProperty(`--body-bg`);
  html.style.removeProperty(`--dark-bg`);

  // reseting to ltr
  ltrFn();

  // reseting to vertical
  verticalFn();
  mainContent.removeEventListener("click", clearNavDropdown);

  // reseting page style
  document.querySelector("#switcher-classic").checked = false;
  document.querySelector("#switcher-regular").checked = true;

  // reseting layout width styles
  document.querySelector("#switcher-full-width").checked = true;
  document.querySelector("#switcher-boxed").checked = false;

  // reseting menu position styles
  document.querySelector("#switcher-menu-fixed").checked = true;
  document.querySelector("#switcher-menu-scroll").checked = false;

  // reseting header position styles
  document.querySelector("#switcher-header-fixed").checked = true;
  document.querySelector("#switcher-header-scroll").checked = false;

  // reseting sidemenu layout styles
  document.querySelector("#switcher-default-menu").checked = true;
  document.querySelector("#switcher-closed-menu").checked = false;
  document.querySelector("#switcher-icontext-menu").checked = false;
  document.querySelector("#switcher-icon-overlay").checked = false;
  document.querySelector("#switcher-detached").checked = false;
  document.querySelector("#switcher-double-menu").checked = false;

  // reseting chart colors
  updateColors();
  document.querySelector(".main-menu").style.marginLeft = "0px";
  document.querySelector(".main-menu").style.marginRight = "0px";

  document.querySelector("#switcher-primary").checked = true;
  document.querySelector("#switcher-background").checked = true;
  document.querySelectorAll(".has-sub").forEach((element) => {
    if (
      element.classList.contains("open") &&
      !element.classList.contains("active")
    ) {
      element.classList.remove("open");
      element.querySelector("ul").style = "";
    }
  });
}

// chart colors
function updateColors() {
  "use strict";
  primaryRGB1 = getComputedStyle(document.documentElement)
    .getPropertyValue("--color-primary-rgb")
    .trim();

  //get variable
  myVarVal = localStorage.getItem("primaryRGB") || primaryRGB1;

  //index
  if (document.querySelector("#salesOverview") !== null) {
    setTimeout(() => {
      salesOverview();
    }, 100);
  }
  if (document.querySelector("#sales-chart") !== null) {
    sparkchart();
  }
  if (document.querySelector("#sales-chart2") !== null) {
    sparkchart2();
  }
  if (document.querySelector("#sales-chart3") !== null) {
    sparkchart3();
  }
  if (document.querySelector("#sales-chart4") !== null) {
    sparkchart4();
  }
  if (document.querySelector("#visitors") !== null) {
    visitorschart();
  }
  if (document.querySelector("#sales-donut") !== null) {
    salesdonut();
  }

  //index-2
  if (document.querySelector("#earnings") !== null) {
    Earnings();
  }

  //index-3
  if (document.querySelector("#crypto") !== null) {
    cryptoCurrency();
  }
  if (document.querySelector("#btc-chart") !== null) {
    crypto();
  }
  if (document.querySelector("#eth-chart") !== null) {
    crypto2();
  }
  if (document.querySelector("#dash-chart") !== null) {
    crypto3();
  }
  if (document.querySelector("#glm-chart") !== null) {
    crypto4();
  }

  //index-4
  if (document.querySelector("#subscriptionOverview") !== null) {
    subOverview();
  }
  if (document.querySelector("#candidates-chart") !== null) {
    Candidates();
  }

  //index-5
  if (document.querySelector("#nft-statistics") !== null) {
    nftStatistics();
  }

  //index-6
  if (document.querySelector("#active-chart") !== null) {
    activechart();
  }
  if (document.querySelector("#audienceReport") !== null) {
    audience();
  }
  if (document.querySelector("#sessions") !== null) {
    Sessions();
  }
  if (document.querySelector("#audience") !== null) {
    audienceOverview();
  }
  if (document.querySelector("#session2") !== null) {
    session2();
  }

  //index-7
  if (document.querySelector("#projectAnalysis") !== null) {
    projectAnalysis();
  }

  //index-8
  if (document.querySelector("#performanceReport") !== null) {
    performanceReport();
  }

  //index-9
  if (document.querySelector("#revenue") !== null) {
    revenueOverview();
  }
  if (document.querySelector("#leads") !== null) {
    leads();
  }

  //index-10
  if (document.querySelector("#statistics") !== null) {
    statistics();
  }

  //index-11
  if (document.querySelector("#totalInvested") !== null) {
    totalInvested();
  }
  if (document.querySelector("#totalInvestmentsStats") !== null) {
    totalInvestmentsStats();
  }

  //index-12
  if (document.querySelector("#earningsReport") !== null) {
    earningsReport();
  }

  //widgets
  if (document.querySelector("#report") !== null) {
    targetReport();
  }
  if (document.querySelector("#views") !== null) {
    pageviews();
  }
}
updateColors();

if (document.querySelector("#hs-overlay-switcher")) {
  /* Horizontal Start */
  if (
    document.querySelector("html").getAttribute("data-nav-layout") ===
      "horizontal" &&
    !document.querySelector(".landing-body") == true
  ) {
    horizontalClickFn();
  }
  /* Horizontal Start */
}
/* RTL Start */
if (document.querySelector("html").getAttribute("dir") === "rtl") {
  if (document.querySelector("#hs-overlay-switcher")) {
    rtlFn();
  }
}

("use strict");

const ANIMATION_DURATION = 300;

const sidebar = document.getElementById("sidebar");
let mainContentDiv = document.querySelector(".main-content");

const slideHasSub = document.querySelectorAll(".nav > ul > .slide.has-sub");

const firstLevelItems = document.querySelectorAll(
  ".nav > ul > .slide.has-sub > a"
);

const innerLevelItems = document.querySelectorAll(
  ".nav > ul > .slide.has-sub .slide.has-sub > a"
);

class PopperObject {
  instance = null;
  reference = null;
  popperTarget = null;

  constructor(reference, popperTarget) {
    this.init(reference, popperTarget);
  }

  init(reference, popperTarget) {
    this.reference = reference;
    this.popperTarget = popperTarget;
    this.instance = Popper.createPopper(this.reference, this.popperTarget, {
      placement: "bottom",
      strategy: "relative",
      resize: true,
      modifiers: [
        {
          name: "computeStyles",
          options: {
            adaptive: false,
          },
        },
      ],
    });

    document.addEventListener(
      "click",
      (e) => this.clicker(e, this.popperTarget, this.reference),
      false
    );

    const ro = new ResizeObserver(() => {
      this.instance.update();
    });

    ro.observe(this.popperTarget);
    ro.observe(this.reference);
  }

  clicker(event, popperTarget, reference) {
    if (
      sidebar.classList.contains("collapsed") &&
      !popperTarget.contains(event.target) &&
      !reference.contains(event.target)
    ) {
      this.hide();
    }
  }

  hide() {}
}

class Poppers {
  subMenuPoppers = [];

  constructor() {
    this.init();
  }

  init() {
    slideHasSub.forEach((element) => {
      this.subMenuPoppers.push(
        new PopperObject(element, element.lastElementChild)
      );
      this.closePoppers();
    });
  }

  togglePopper(target) {
    if (
      window.getComputedStyle(target).visibility === "hidden" &&
      window.getComputedStyle(target).visibility === undefined
    )
      target.style.visibility = "visible";
    else target.style.visibility = "hidden";
  }

  updatePoppers() {
    this.subMenuPoppers.forEach((element) => {
      element.instance.state.elements.popper.style.display = "none";
      element.instance.update();
    });
  }

  closePoppers() {
    this.subMenuPoppers.forEach((element) => {
      element.hide();
    });
  }
}

const slideUp = (target, duration = ANIMATION_DURATION) => {
  const { parentElement } = target;
  parentElement.classList.remove("open");
  target.style.transitionProperty = "height, margin, padding";
  target.style.transitionDuration = `${duration}ms`;
  target.style.boxSizing = "border-box";
  target.style.height = `${target.offsetHeight}px`;
  target.offsetHeight;
  target.style.overflow = "hidden";
  target.style.height = 0;
  target.style.paddingTop = 0;
  target.style.paddingBottom = 0;
  target.style.marginTop = 0;
  target.style.marginBottom = 0;
  window.setTimeout(() => {
    target.style.display = "none";
    target.style.removeProperty("height");
    target.style.removeProperty("padding-top");
    target.style.removeProperty("padding-bottom");
    target.style.removeProperty("margin-top");
    target.style.removeProperty("margin-bottom");
    target.style.removeProperty("overflow");
    target.style.removeProperty("transition-duration");
    target.style.removeProperty("transition-property");
  }, duration);
};
const slideDown = (target, duration = ANIMATION_DURATION) => {
  const { parentElement } = target;
  parentElement.classList.add("open");
  target.style.removeProperty("display");
  let { display } = window.getComputedStyle(target);
  if (display === "none") display = "block";
  target.style.display = display;
  const height = target.offsetHeight;
  target.style.overflow = "hidden";
  target.style.height = 0;
  target.style.paddingTop = 0;
  target.style.paddingBottom = 0;
  target.style.marginTop = 0;
  target.style.marginBottom = 0;
  target.offsetHeight;
  target.style.boxSizing = "border-box";
  target.style.transitionProperty = "height, margin, padding";
  target.style.transitionDuration = `${duration}ms`;
  target.style.height = `${height}px`;
  target.style.removeProperty("padding-top");
  target.style.removeProperty("padding-bottom");
  target.style.removeProperty("margin-top");
  target.style.removeProperty("margin-bottom");
  window.setTimeout(() => {
    target.style.removeProperty("height");
    target.style.removeProperty("overflow");
    target.style.removeProperty("transition-duration");
    target.style.removeProperty("transition-property");
  }, duration);
};
const slideToggle = (target, duration = ANIMATION_DURATION) => {
  let html = document.querySelector("html");
  if (
    !(
      (html.getAttribute("data-nav-style") === "menu-hover" &&
        html.getAttribute("toggled") === "menu-hover-closed" &&
        window.innerWidth >= 992) ||
      (html.getAttribute("data-nav-style") === "icon-hover" &&
        html.getAttribute("toggled") === "icon-hover-closed" &&
        window.innerWidth >= 992)
    ) &&
    target &&
    target.nodeType != 3
  ) {
    if (window.getComputedStyle(target).display === "none") {
      return slideDown(target, duration);
    }
    return slideUp(target, duration);
  }
};

const PoppersInstance = new Poppers();

/**
 * wait for the current animation to finish and update poppers position
 */
const updatePoppersTimeout = () => {
  setTimeout(() => {
    PoppersInstance.updatePoppers();
  }, ANIMATION_DURATION);
};

const defaultOpenMenus = document.querySelectorAll(".slide.has-sub.open");

defaultOpenMenus.forEach((element) => {
  element.lastElementChild.style.display = "block";
});

/**
 * handle top level submenu click
 */
firstLevelItems.forEach((element) => {
  element.addEventListener("click", () => {
    let html = document.querySelector("html");
    if (
      !(
        (html.getAttribute("data-nav-style") === "menu-hover" &&
          html.getAttribute("toggled") === "menu-hover-closed" &&
          window.innerWidth >= 992) ||
        (html.getAttribute("data-nav-style") === "icon-hover" &&
          html.getAttribute("toggled") === "icon-hover-closed" &&
          window.innerWidth >= 992)
      )
    ) {
      const parentMenu = element.closest(".nav.sub-open");
      if (parentMenu)
        parentMenu
          .querySelectorAll(":scope > ul > .slide.has-sub > a")
          .forEach((el) => {
            if (
              el.nextElementSibling.style.display === "block" ||
              window.getComputedStyle(el.nextElementSibling).display === "block"
            ) {
              slideUp(el.nextElementSibling);
            }
          });
      slideToggle(element.nextElementSibling);
    }
  });
});

/**
 * handle inner submenu click
 */
innerLevelItems.forEach((element) => {
  let html = document.querySelector("html");
  element.addEventListener("click", () => {
    const innerMenu = element.closest(".slide-menu");
    if (innerMenu)
      innerMenu.querySelectorAll(":scope .slide.has-sub > a").forEach((el) => {
        if (
          el.nextElementSibling &&
          el.nextElementSibling?.style.display === "block"
        ) {
          slideUp(el.nextElementSibling);
        }
      });
    slideToggle(element.nextElementSibling);
  });
});

/**
 * menu styles
 */

window.addEventListener("resize", () => {
  let mainContent = document.querySelector(".main-content");
  setTimeout(() => {
    if (window.innerWidth <= 992) {
      mainContent.addEventListener("click", menuClose);
    } else {
      mainContent.removeEventListener("click", menuClose);
    }
  }, 100);
});
let headerToggleBtn, WindowPreSize;
(() => {
  let html = document.querySelector("html");
  headerToggleBtn = document.querySelector(".sidebar-toggle");
  headerToggleBtn.addEventListener("click", toggleSidemenu);
  let mainContent = document.querySelector(".main-content");
  if (window.innerWidth <= 992) {
    mainContent.addEventListener("click", menuClose);
  } else {
    mainContent.removeEventListener("click", menuClose);
  }

  WindowPreSize = [window.innerWidth];
  setNavActive();
  if (
    html.getAttribute("data-nav-layout") === "horizontal" &&
    window.innerWidth >= 992
  ) {
    clearNavDropdown();
    mainContent.addEventListener("click", clearNavDropdown);
  } else {
    mainContent.removeEventListener("click", clearNavDropdown);
  }

  window.addEventListener("resize", ResizeMenu);
  switcherArrowFn();

  if (
    !localStorage.getItem("Syntolayout") &&
    !localStorage.getItem("Syntonavstyles") &&
    !localStorage.getItem("Syntoverticalstyles") &&
    !document.querySelector(".landing-body") &&
    document.querySelector("html").getAttribute("data-nav-layout") !==
      "horizontal"
  ) {
    // To enable sidemenu layout styles
    // iconTextFn()
    // detachedFn();
    // closedSidemenuFn();
    // doubletFn();
    if (
      !html.getAttribute("data-vertical-style") &&
      !html.getAttribute("data-nav-style")
    ) {
      iconOverayFn();
    }
  }

  toggleSidemenu();

  if (
    (html.getAttribute("data-nav-style") === "menu-hover" ||
      html.getAttribute("data-nav-style") === "icon-hover") &&
    window.innerWidth >= 992
  ) {
    clearNavDropdown();
  }
  if (window.innerWidth < 992) {
    html.setAttribute("toggled", "close");
  }
})();

function ResizeMenu() {
  let html = document.querySelector("html");
  let mainContent = document.querySelector(".main-content");

  WindowPreSize.push(window.innerWidth);
  if (WindowPreSize.length > 2) {
    WindowPreSize.shift();
  }
  if (WindowPreSize.length > 1) {
    if (
      WindowPreSize[WindowPreSize.length - 1] < 992 &&
      WindowPreSize[WindowPreSize.length - 2] >= 992
    ) {
      // less than 992;
      mainContent.addEventListener("click", menuClose);
      setNavActive();
      toggleSidemenu();
      mainContent.removeEventListener("click", clearNavDropdown);
    }

    if (
      WindowPreSize[WindowPreSize.length - 1] >= 992 &&
      WindowPreSize[WindowPreSize.length - 2] < 992
    ) {
      // greater than 992
      mainContent.removeEventListener("click", menuClose);
      toggleSidemenu();
      if (html.getAttribute("data-nav-layout") === "horizontal") {
        clearNavDropdown();
        mainContent.addEventListener("click", clearNavDropdown);
      } else {
        mainContent.removeEventListener("click", clearNavDropdown);
      }
      if (
        !document.querySelector("html").getAttribute("toggled") ==
        "double-menu-open"
      ) {
        html.removeAttribute("toggled");
      }
    }
  }
  checkHoriMenu();
}
function menuClose() {
  let html = document.querySelector("html");
  html.setAttribute("toggled", "close");
  document.querySelector("#responsive-overlay").classList.remove("active");
}
function toggleSidemenu() {
  let html = document.querySelector("html");
  let sidemenuType = html.getAttribute("data-nav-layout");

  if (window.innerWidth >= 992) {
    if (sidemenuType === "vertical") {
      sidebar.removeEventListener("mouseenter", mouseEntered);
      sidebar.removeEventListener("mouseleave", mouseLeave);
      sidebar.removeEventListener("click", icontextOpen);
      mainContentDiv.removeEventListener("click", icontextClose);
      let sidemenulink = document.querySelectorAll(
        ".main-menu li > .side-menu__item"
      );
      sidemenulink.forEach((ele) =>
        ele.removeEventListener("click", doubleClickFn)
      );

      let verticalStyle = html.getAttribute("data-vertical-style");
      switch (verticalStyle) {
        // closed
        case "closed":
          html.removeAttribute("data-nav-style");
          if (html.getAttribute("toggled") === "close-menu-close") {
            html.removeAttribute("toggled");
          } else {
            html.setAttribute("toggled", "close-menu-close");
          }
          break;
        // icon-overlay
        case "overlay":
          html.removeAttribute("data-nav-style");
          if (html.getAttribute("toggled") === "icon-overlay-close") {
            html.removeAttribute("toggled", "icon-overlay-close");
            sidebar.removeEventListener("mouseenter", mouseEntered);
            sidebar.removeEventListener("mouseleave", mouseLeave);
          } else {
            if (window.innerWidth >= 992) {
              html.setAttribute("toggled", "icon-overlay-close");
              sidebar.addEventListener("mouseenter", mouseEntered);
              sidebar.addEventListener("mouseleave", mouseLeave);
            } else {
              sidebar.removeEventListener("mouseenter", mouseEntered);
              sidebar.removeEventListener("mouseleave", mouseLeave);
            }
          }
          break;
        // icon-text
        case "icontext":
          html.removeAttribute("data-nav-style");
          if (html.getAttribute("toggled") === "icon-text-close") {
            html.removeAttribute("toggled", "icon-text-close");
            sidebar.removeEventListener("click", icontextOpen);
            mainContentDiv.removeEventListener("click", icontextClose);
          } else {
            html.setAttribute("toggled", "icon-text-close");
            if (window.innerWidth >= 992) {
              sidebar.addEventListener("click", icontextOpen);
              mainContentDiv.addEventListener("click", icontextClose);
            } else {
              sidebar.removeEventListener("click", icontextOpen);
              mainContentDiv.removeEventListener("click", icontextClose);
            }
          }
          break;
        // doublemenu
        case "doublemenu":
          html.removeAttribute("data-nav-style");
          if (html.getAttribute("toggled") === "double-menu-open") {
            html.setAttribute("toggled", "double-menu-close");
            if (document.querySelector(".slide-menu")) {
              let slidemenu = document.querySelectorAll(".slide-menu");
              slidemenu.forEach((e) => {
                if (e.classList.contains("double-menu-active")) {
                  e.classList.remove("double-menu-active");
                }
              });
            }
          } else {
            let sidemenu = document.querySelector(".side-menu__item.active");
            if (sidemenu) {
              html.setAttribute("toggled", "double-menu-open");
              if (sidemenu.nextElementSibling) {
                sidemenu.nextElementSibling.classList.add("double-menu-active");
              } else {
                document.querySelector("html").setAttribute("toggled", "");
              }
            }
          }

          doublemenu();
          break;
        // detached
        case "detached":
          if (html.getAttribute("toggled") === "detached-close") {
            html.removeAttribute("toggled", "detached-close");
            sidebar.removeEventListener("mouseenter", mouseEntered);
            sidebar.removeEventListener("mouseleave", mouseLeave);
          } else {
            html.setAttribute("toggled", "detached-close");
            if (window.innerWidth >= 992) {
              sidebar.addEventListener("mouseenter", mouseEntered);
              sidebar.addEventListener("mouseleave", mouseLeave);
            } else {
              sidebar.removeEventListener("mouseenter", mouseEntered);
              sidebar.removeEventListener("mouseleave", mouseLeave);
            }
          }
          break;
        // defaultlo
        case "default":
          iconOverayFn();
          html.removeAttribute("toggled");
      }
      let menuclickStyle = html.getAttribute("data-nav-style");
      switch (menuclickStyle) {
        case "menu-click":
          if (html.getAttribute("toggled") === "menu-click-closed") {
            html.removeAttribute("toggled");
          } else {
            html.setAttribute("toggled", "menu-click-closed");
          }
          break;
        case "menu-hover":
          if (html.getAttribute("toggled") === "menu-hover-closed") {
            html.removeAttribute("toggled");
            setNavActive();
          } else {
            html.setAttribute("toggled", "menu-hover-closed");
            clearNavDropdown();
          }
          break;
        case "icon-click":
          if (html.getAttribute("toggled") === "icon-click-closed") {
            html.removeAttribute("toggled");
          } else {
            html.setAttribute("toggled", "icon-click-closed");
          }
          break;
        case "icon-hover":
          if (html.getAttribute("toggled") === "icon-hover-closed") {
            html.removeAttribute("toggled");
            setNavActive();
          } else {
            html.setAttribute("toggled", "icon-hover-closed");
            clearNavDropdown();
          }
          break;

        //for making any horizontal style as default
        default:
      }
    }
  } else {
    if (html.getAttribute("toggled") === "close") {
      html.setAttribute("toggled", "open");
      let i = document.createElement("div");
      i.id = "responsive-overlay";
      setTimeout(() => {
        if (document.querySelector("html").getAttribute("toggled") == "open") {
          document.querySelector("#responsive-overlay").classList.add("active");
          document
            .querySelector("#responsive-overlay")
            .addEventListener("click", () => {
              document
                .querySelector("#responsive-overlay")
                .classList.remove("active");
              menuClose();
            });
        }
        window.addEventListener("resize", () => {
          if (window.screen.width >= 992) {
            document
              .querySelector("#responsive-overlay")
              .classList.remove("active");
          }
        });
      }, 100);
    } else {
      html.setAttribute("toggled", "close");
    }
  }
}
function mouseEntered() {
  let html = document.querySelector("html");
  html.setAttribute("icon-overlay", "open");
}
function mouseLeave() {
  let html = document.querySelector("html");
  html.removeAttribute("icon-overlay");
}
function icontextOpen() {
  let html = document.querySelector("html");
  html.setAttribute("icon-text", "open");
}
function icontextClose() {
  let html = document.querySelector("html");
  html.removeAttribute("icon-text");
}
function closedSidemenuFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-layout", "vertical");
  html.setAttribute("data-vertical-style", "closed");
  toggleSidemenu();
}
function detachedFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-layout", "vertical");
  html.setAttribute("data-vertical-style", "detached");
  toggleSidemenu();
}
function iconTextFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-layout", "vertical");
  html.setAttribute("data-vertical-style", "icontext");
  toggleSidemenu();
}
function iconOverayFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-layout", "vertical");
  html.setAttribute("data-vertical-style", "overlay");
  toggleSidemenu();
}
function doubletFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-layout", "vertical");
  html.setAttribute("data-vertical-style", "doublemenu");
  toggleSidemenu();

  // Select the menu slide item
  const menuSlideItem = document.querySelectorAll(
    ".main-menu > li > .side-menu__item"
  );

  // Create the tooltip element
  const tooltip = document.createElement("div");
  // tooltip.textContent = "This is a tooltip";

  // Set the CSS properties of the tooltip element
  tooltip.style.setProperty("position", "fixed");
  tooltip.style.setProperty("display", "none");
  tooltip.style.setProperty("padding", "0.5rem");
  tooltip.style.setProperty("font-weight", "500");
  tooltip.style.setProperty("font-size", "0.75rem");
  tooltip.style.setProperty("background-color", "rgb(15, 23 ,42)");
  tooltip.style.setProperty("color", "rgb(255, 255 ,255)");
  tooltip.style.setProperty("margin-inline-start", "45px");
  tooltip.style.setProperty("border-radius", "0.25rem");
  tooltip.style.setProperty("z-index", "99");

  menuSlideItem.forEach((e) => {
    // Add an event listener to the menu slide item to show the tooltip
    e.addEventListener("mouseenter", () => {
      tooltip.style.setProperty("display", "block");
      tooltip.textContent = e.querySelector(".side-menu__label").textContent;
      if (
        document.querySelector("html").getAttribute("data-vertical-style") ==
        "doublemenu"
      ) {
        e.appendChild(tooltip);
      }
    });

    // Add an event listener to hide the tooltip
    e.addEventListener("mouseleave", () => {
      tooltip.style.setProperty("display", "none");
      tooltip.textContent = e.querySelector(".side-menu__label").textContent;
      if (
        document.querySelector("html").getAttribute("data-vertical-style") ==
        "doublemenu"
      ) {
        e.removeChild(tooltip);
      }
    });
  });
}
function menuClickFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-style", "menu-click");

  toggleSidemenu();
  if (html.getAttribute("data-nav-layout") === "vertical") {
    setNavActive();
  }
}
function menuhoverFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-style", "menu-hover");
  html.removeAttribute("hor-style");
  html.removeAttribute("data-vertical-style");
  toggleSidemenu();
  if (html.getAttribute("toggled") === "menu-hover-closed") {
    clearNavDropdown();
  }
}
function iconClickFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-style", "icon-click");
  toggleSidemenu();
  if (html.getAttribute("data-nav-layout") === "vertical") {
    setNavActive();
  }
}
function iconHoverFn() {
  let html = document.querySelector("html");
  html.setAttribute("data-nav-style", "icon-hover");
  toggleSidemenu();
  clearNavDropdown();
}
function setNavActive() {
  let currentPath = window.location.pathname.split("/")[0];

  currentPath =
    location.pathname == "/" ? "index.html" : location.pathname.substring(1);
  currentPath = currentPath.substring(currentPath.lastIndexOf("/") + 1);
  let sidemenuItems = document.querySelectorAll(".side-menu__item");
  sidemenuItems.forEach((e) => {
    if (currentPath === "/") {
      currentPath = "index.html";
    }
    if (e.getAttribute("href") === currentPath) {
      e.classList.add("active");
      e.parentElement.classList.add("active");
      let parent = e.closest("ul");
      let parentNotMain = e.closest("ul:not(.main-menu)");
      let hasParent = true;
      while (hasParent) {
        if (parent) {
          parent.classList.add("active");
          parent.previousElementSibling.classList.add("active");
          parent.parentElement.classList.add("active");
          slideDown(parent);
          parent = parent.parentElement.closest("ul");
          //
          if (parent && parent.closest("ul:not(.main-menu)")) {
            parentNotMain = parent.closest("ul:not(.main-menu)");
          }
          if (!parentNotMain) hasParent = false;
        } else {
          hasParent = false;
        }
      }
    }
  });
}
function clearNavDropdown() {
  let sidemenus = document.querySelectorAll("ul.slide-menu");
  sidemenus.forEach((e) => {
    let parent = e.closest("ul");
    let parentNotMain = e.closest("ul:not(.main-menu)");
    if (parent) {
      let hasParent = parent.closest("ul.main-menu");
      while (hasParent) {
        parent.classList.add("active");
        slideUp(parent);
        //
        parent = parent.parentElement.closest("ul");
        parentNotMain = parent.closest("ul:not(.main-menu)");
        if (!parentNotMain) hasParent = false;
      }
    }
  });
}
function switcherArrowFn() {
  // used to remove is-expanded class and remove class on clicking arrow buttons
  function slideClick() {
    let slide = document.querySelectorAll(".slide");
    let slideMenu = document.querySelectorAll(".slide-menu");
    slide.forEach((element, index) => {
      if (element.classList.contains("is-expanded") == true) {
        element.classList.remove("is-expanded");
      }
    });
    slideMenu.forEach((element, index) => {
      if (element.classList.contains("open") == true) {
        element.classList.remove("open");
        element.style.display = "none";
      }
    });
  }

  slideClick();
}
let slideLeft = document.querySelector(".slide-left");
let slideRight = document.querySelector(".slide-right");
slideLeft.addEventListener("click", () => {
  let menuNav = document.querySelector(".main-menu");
  let mainContainer1 = document.querySelector(".main-sidebar");
  let marginLeftValue = Math.ceil(
    Number(window.getComputedStyle(menuNav).marginLeft.split("px")[0])
  );
  let marginRightValue = Math.ceil(
    Number(window.getComputedStyle(menuNav).marginRight.split("px")[0])
  );
  let mainContainer1Width = mainContainer1.offsetWidth;
  if (menuNav.scrollWidth > mainContainer1.offsetWidth) {
    if (!(document.querySelector("html").getAttribute("dir") === "rtl")) {
      if (
        marginLeftValue < 0 &&
        !(Math.abs(marginLeftValue) < mainContainer1Width)
      ) {
        menuNav.style.marginRight = 0;
        menuNav.style.marginLeft =
          Number(menuNav.style.marginLeft.split("px")[0]) +
          Math.abs(mainContainer1Width) +
          "px";
        slideRight.classList.remove("hidden");
      } else if (marginLeftValue >= 0) {
        menuNav.style.marginLeft = "0px";
        slideLeft.classList.add("hidden");
        slideRight.classList.remove("hidden");
      } else {
        menuNav.style.marginLeft = "0px";
        slideLeft.classList.add("hidden");
        slideRight.classList.remove("hidden");
      }
    } else {
      if (
        marginRightValue < 0 &&
        !(Math.abs(marginRightValue) < mainContainer1Width)
      ) {
        menuNav.style.marginLeft = 0;
        menuNav.style.marginRight =
          Number(menuNav.style.marginRight.split("px")[0]) +
          Math.abs(mainContainer1Width) +
          "px";
        slideRight.classList.remove("hidden");
      } else if (marginRightValue >= 0) {
        menuNav.style.marginRight = "0px";
        slideLeft.classList.add("hidden");
        slideRight.classList.remove("hidden");
      } else {
        menuNav.style.marginRight = "0px";
        slideLeft.classList.add("hidden");
        slideRight.classList.remove("hidden");
      }
    }
  } else {
    document.querySelector(".main-menu").style.marginLeft = "0px";
    document.querySelector(".main-menu").style.marginRight = "0px";
  }
  let element = document.querySelector(".main-menu > .slide.open");
  let element1 = document.querySelector(".main-menu > .slide.open >ul");
  element.classList.remove("active");
  element1.style.display = "none";

  switcherArrowFn();
  return;
  //
});
slideRight.addEventListener("click", () => {
  let menuNav = document.querySelector(".main-menu");
  let mainContainer1 = document.querySelector(".main-sidebar");
  let marginLeftValue = Math.ceil(
    Number(window.getComputedStyle(menuNav).marginLeft.split("px")[0])
  );
  let marginRightValue = Math.ceil(
    Number(window.getComputedStyle(menuNav).marginRight.split("px")[0])
  );
  let check = menuNav.scrollWidth - mainContainer1.offsetWidth;
  let mainContainer1Width = mainContainer1.offsetWidth;

  if (menuNav.scrollWidth > mainContainer1.offsetWidth) {
    if (!(document.querySelector("html").getAttribute("dir") === "rtl")) {
      if (Math.abs(check) > Math.abs(marginLeftValue)) {
        menuNav.style.marginRight = 0;
        if (
          !(Math.abs(check) > Math.abs(marginLeftValue) + mainContainer1Width)
        ) {
          mainContainer1Width = Math.abs(check) - Math.abs(marginLeftValue);
          slideRight.classList.add("hidden");
        }
        menuNav.style.marginLeft =
          Number(menuNav.style.marginLeft.split("px")[0]) -
          Math.abs(mainContainer1Width) +
          "px";
        slideLeft.classList.remove("hidden");
      }
    } else {
      if (Math.abs(check) > Math.abs(marginRightValue)) {
        menuNav.style.marginLeft = 0;
        if (
          !(Math.abs(check) > Math.abs(marginRightValue) + mainContainer1Width)
        ) {
          mainContainer1Width = Math.abs(check) - Math.abs(marginRightValue);
          slideRight.classList.add("hidden");
        }
        menuNav.style.marginRight =
          Number(menuNav.style.marginRight.split("px")[0]) -
          Math.abs(mainContainer1Width) +
          "px";
        slideLeft.classList.remove("hidden");
      }
    }
  }
  let element = document.querySelector(".main-menu > .slide.open");
  let element1 = document.querySelector(".main-menu > .slide.open >ul");
  element.classList.remove("active");
  element1.style.display = "none";

  switcherArrowFn();
  return;
});
function checkHoriMenu() {
  let menuNav = document.querySelector(".main-menu");
  let mainContainer1 = document.querySelector(".main-sidebar");
  let slideLeft = document.querySelector(".slide-left");
  let slideRight = document.querySelector(".slide-right");
  let marginLeftValue = Math.ceil(
    Number(window.getComputedStyle(menuNav).marginLeft.split("px")[0])
  );
  let marginRightValue = Math.ceil(
    Number(window.getComputedStyle(menuNav).marginRight.split("px")[0])
  );
  let check = menuNav.scrollWidth - mainContainer1.offsetWidth;
  // Show/Hide the arrows
  if (menuNav.scrollWidth > mainContainer1.offsetWidth) {
    slideRight.classList.remove("hidden");
    slideLeft.classList.add("hidden");
  } else {
    slideRight.classList.add("hidden");
    slideLeft.classList.add("hidden");
    menuNav.style.marginLeft = "0px";
    menuNav.style.marginRight = "0px";
  }
  if (!(document.querySelector("html").getAttribute("dir") === "rtl")) {
    // LTR check the width and adjust the menu in screen
    if (menuNav.scrollWidth > mainContainer1.offsetWidth) {
      if (Math.abs(check) < Math.abs(marginLeftValue)) {
        menuNav.style.marginLeft = -check + "px";
        slideLeft.classList.remove("hidden");
        slideRight.classList.add("hidden");
      }
    }
    if (marginLeftValue == 0) {
      slideLeft.classList.add("hidden");
      slideRight.classList.remove("hidden");
    } else {
      slideLeft.classList.remove("hidden");
    }
  } else {
    // RTL check the width and adjust the menu in screen
    if (menuNav.scrollWidth > mainContainer1.offsetWidth) {
      if (Math.abs(check) < Math.abs(marginRightValue)) {
        menuNav.style.marginRight = -check + "px";
        slideLeft.classList.remove("hidden");
        slideRight.classList.add("hidden");
      }
    }
    if (marginRightValue == 0) {
      slideLeft.classList.add("hidden");
    } else {
      slideLeft.classList.remove("hidden");
    }
  }
  if (marginLeftValue != 0 || marginRightValue != 0) {
    slideLeft.classList.remove("hidden");
  }
}

// double-menu click toggle start
function doublemenu() {
  if (window.innerWidth >= 992) {
    let html = document.querySelector("html");
    let sidemenulink = document.querySelectorAll(
      ".main-menu > li > .side-menu__item"
    );
    sidemenulink.forEach((ele) => {
      ele.addEventListener("click", doubleClickFn);
    });
  }
}
function doubleClickFn() {
  var $this = this;
  let html = document.querySelector("html");
  var checkElement = $this.nextElementSibling;
  if (checkElement) {
    if (!checkElement.classList.contains("double-menu-active")) {
      if (document.querySelector(".slide-menu")) {
        let slidemenu = document.querySelectorAll(".slide-menu");
        slidemenu.forEach((e) => {
          if (e.classList.contains("double-menu-active")) {
            e.classList.remove("double-menu-active");
            html.setAttribute("toggled", "double-menu-close");
          }
        });
      }
      checkElement.classList.add("double-menu-active");
      html.setAttribute("toggled", "double-menu-open");
    }
  }
}
// double-menu click toggle end

window.addEventListener("unload", () => {
  let mainContent = document.querySelector(".main-content");
  mainContent.removeEventListener("click", clearNavDropdown);
  window.removeEventListener("resize", ResizeMenu);
  let sidemenulink = document.querySelectorAll(
    ".main-menu li > .side-menu__item"
  );
  sidemenulink.forEach((ele) =>
    ele.removeEventListener("click", doubleClickFn)
  );
});

let customScrollTop = () => {
  document.querySelectorAll(".side-menu__item").forEach((ele) => {
    if (ele.classList.contains("active")) {
      let elemRect = ele.getBoundingClientRect();
      if (
        ele.children[0] &&
        ele.parentElement.classList.contains("has-sub") &&
        elemRect.top > 435
      ) {
        ele.scrollIntoView({ behavior: "smooth" });
      }
    }
  });
};
setTimeout(() => {
  customScrollTop();
}, 300);

// default menu
// For Horizontal menu Overflow
document.querySelectorAll(".side-menu__item").forEach((element) => {
  element.addEventListener("click", () => {
    let ulMenu = element.parentNode.querySelector(".child2");
    if (ulMenu) {
      const elementRect = ulMenu.getBoundingClientRect();
      const isOverflowing =
        elementRect.right > window.innerWidth ||
        elementRect.bottom > window.innerHeight;
      if (
        isOverflowing &&
        document.querySelector("html").getAttribute("data-nav-layout") ==
          "horizontal" &&
        document.querySelector("html").getAttribute("dir") == "ltr"
      ) {
        ulMenu.style.setProperty("right", "100%", "important");
        ulMenu.style.setProperty("left", "auto", "important");
      }

      const isOverflowingRTL = ulMenu.scrollWidth <= ulMenu.clientWidth;
      if (
        isOverflowingRTL &&
        ulMenu.scrollWidth != 0 &&
        document.querySelector("html").getAttribute("data-nav-layout") ==
          "horizontal" &&
        document.querySelector("html").getAttribute("dir") == "rtl"
      ) {
        ulMenu.style.setProperty("left", "100%", "important");
        ulMenu.style.setProperty("right", "auto", "important");
      }
    }
  });
});

// sticky
("use strict");
(() => {
  var navbar = document.querySelector(".header");
  var navbar1 = document.querySelector(".app-sidebar");
  var sticky = navbar.clientHeight;
  var sticky1 = navbar1.clientHeight;
  function stickyFn() {
    if (window.pageYOffset >= sticky) {
      navbar.classList.add("sticky-pin");
      navbar1.classList.add("sticky-pin");
    } else {
      navbar.classList.remove("sticky-pin");
      navbar1.classList.remove("sticky-pin");
    }
  }
  window.addEventListener("scroll", stickyFn);
  window.addEventListener("DOMContentLoaded", stickyFn);
})();

// switch
window.addEventListener("load", () => {
  const themeBtn = document.querySelectorAll("[data-hs-theme-click-value]");
  let html = document.querySelector("html");

  themeBtn.forEach(($item) => {
    $item.addEventListener("click", () => {
      if (html.classList.contains("dark")) {
        html.classList.remove("dark");
        localStorage.removeItem("layout-theme");
        localStorage.removeItem("Syntodarktheme");
        localStorage.removeItem("SyntoMenu");
        localStorage.removeItem("SyntoHeader");
        localStorage.removeItem("darkBgRGB");
        localStorage.removeItem("bodyBgRGB");
        html.setAttribute("data-menu-styles", "dark");
        html.setAttribute("data-header-styles", "light");
        html.style.removeProperty("--body-bg");
        html.style.removeProperty("--dark-bg");
        if (document.querySelector("#hs-overlay-switcher")) {
          document.getElementById("switcher-light-theme").checked = true;
        }
      } else {
        if (document.querySelector("#hs-overlay-switcher")) {
          document.getElementById("switcher-dark-theme").checked = true;
        }
        html.setAttribute("class", "dark");
        localStorage.setItem("layout-theme", "dark");
        html.setAttribute("data-header-styles", "dark");
        html.setAttribute("data-menu-styles", "dark");
        localStorage.removeItem("SyntoMenu");
        localStorage.removeItem("SyntoHeader");
        localStorage.setItem("Syntodarktheme", true);
        localStorage.removeItem("Syntolighttheme");
        localStorage.setItem("SyntoMenu", "dark");
        localStorage.setItem("SyntoHeader", "dark");
      }
    });
  });
});
