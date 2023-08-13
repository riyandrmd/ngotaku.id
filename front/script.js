
fetch('http://localhost:5000/getanimefull')
.then((response) => response.json())
.then((data) => {
    
let a = 0
  for (let x of data) {
    const ongoing = document.querySelector(".ongoing")
    const upcoming = document.querySelector(".upcoming")
    const videotrailer = document.querySelector(".video-container")
    const video = document.createElement("iframe")
    const vidcards = document.createElement("div")

    const cards = document.createElement("div")
    const poster = document.createElement("img")
    const title1 = document.createElement("h4")
    const title2 = document.createElement("div")
    const rate = document.createElement("span")
    const images = document.createElement("div")
    const genre = document.createElement("i")
    const studio = document.createElement("p")


    video.src = data[a].trailer
    video.setAttribute("allow","fullscreen")
    vidcards.classList.add("video")
    vidcards.title = data[a].title;
    title2.innerText = data[a].title;
    title2.classList.add("title")

    cards.title = data[a].title;
    title1.classList.add("title")
    poster.src = data[a].poster_url;
    rate.innerText = data[a].rating;
    genre.innerText = data[a].genre + " ";
    studio.innerText = data[a].Studio_Name + "~";
    title1.innerText = data[a].title;

    cards.classList.add("card")
    images.classList.add("images")
    images.appendChild(poster);
    cards.appendChild(images);
    cards.appendChild(title1);
    cards.appendChild(genre);
    cards.appendChild(rate);
    cards.appendChild(studio);

    vidcards.appendChild(video)
    vidcards.appendChild(title2)
    videotrailer.appendChild(vidcards)
    
    if(data[a].status == "ongoing"){
      ongoing.appendChild(cards);
    }
    else {       
      upcoming.appendChild(cards);
    }
    a++;
  }
})


function scrolling(buff){
    buff.addEventListener('wheel',(e)=> {
        e.preventDefault();
        buff.scrollLeft += e.deltaX;
        buff.scrollLeft += e.deltaY;
    })
}

const scrollUpcoming = document.querySelector(".upcoming");
const scrollOngoing = document.querySelector(".ongoing");
const scrollRecommend = document.querySelector(".recommendation")
const scrollVideo = document.querySelector(".video-container")
scrolling(scrollOngoing);
scrolling(scrollRecommend);
scrolling(scrollUpcoming)
scrolling(scrollVideo)


function menuDrop() {
    document.getElementById("myDropdown").classList.toggle("show");
}
  
window.onclick = function(event) {
    if (!event.target.matches('#menu')) {
      var dropdowns = document.getElementsByClassName("dropdown-content");
      var i;
      for (i = 0; i < dropdowns.length; i++) {
        var openDropdown = dropdowns[i];
        if (openDropdown.classList.contains('show')) {
          openDropdown.classList.remove('show');
        }
      }
    }
}
  