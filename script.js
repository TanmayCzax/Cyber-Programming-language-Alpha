function goToRepo() {
  // 🔴 CHANGE THIS URL TO YOUR ACTUAL GITHUB REPO
  window.location.href = "https://github.com/YOUR_USERNAME/YOUR_REPO";
}

function scrollToSection(id) {
  document.getElementById(id).scrollIntoView({ behavior: "smooth" });
}

document.getElementById("year").textContent = new Date().getFullYear();
