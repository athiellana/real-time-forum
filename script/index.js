//let socket = new WebSocket("ws://127.0.0.1:8080/ws");
let socket = new WebSocket("ws://localhost:8080/ws")
let connected = false

        console.log("Attempting Connection...");

        socket.onopen = () => {
            console.log("Successfully Connected");
            socket.send("Hi From the Client!")

            if (document.cookie.split(';').some((item) => item.trim().startsWith('account='))) {
              connected = true
            }

            switchOverlay(connected)
        };

        socket.addEventListener("message", (event) => {
          var data = event.data
          
          
          if (isJsonString(data)) {
            var content = JSON.parse(data)

            if (typeof content == "object") {
              if (content.Message == "cookie") {
                  createCookie("account", content.Value)
                  connected = true
              }
            } 
          }

          switchOverlay(connected)
        })
        
        
        socket.onclose = event => {
          console.log("Socket Closed Connection: ", event);
          socket.send("Client Closed!")
        };
        
        socket.onerror = error => {
          console.log("Socket Error: ", error);
        };
        
        
        // RECUPERATION DES DONNEES DU REGISTER
        
function register(form) {
    const info = {
    "message":"register",
    "firstName": form.elements.first_name.value,
    "lastName": form.elements.last_name.value,
    "username": form.elements.username.value,
    "gender": form.elements.gender.value,
    "email": form.elements.email.value,
    "password": form.elements.password.value,
    "repeatPassword": form.elements.password.value,
    "age": form.elements.age.value
    }
    console.log(info)
    socket.send(JSON.stringify(info))
}

// RECUPERATION DES DONNEES DU LOGIN

function login(form) {
  const identification = {
    "message":"login",
    "usernameMail" : form.elements.username_mail.value,
    "passwordLogin" : form.elements.password_login.value
  }
  console.log(identification)
  socket.send(JSON.stringify(identification))
}

function post(form) {
  const content = {
    "message":"post",
    "contentPost" : form.elements.create_post_content.value
  }
  console.log(content)
}



//GESTION DES DIVS POUR LE ONE PAGE

//VARIABLES
let registerForm = document.getElementById("register_form")
let loginRegisterButton = document.getElementById("login_register")
let loginForm = document.getElementById("login_form")
let registerLoginButton = document.getElementById("register_login")

loginRegisterButton.addEventListener("click", () => {
    if(getComputedStyle(registerForm).display != "none"){
      registerForm.style.display = "none";
      loginForm.style.display="block";
  } else {
    registerForm.style.display = "block";
  }
})

registerLoginButton.addEventListener("click", () => {
    if(getComputedStyle(loginForm).display != "none"){
      loginForm.style.display = "none";
      registerForm.style.display="block";
  } else {
    loginForm.style.display = "block";
  }
})

// c'est thomas qui a mis ça ça marche il est trop fort <3
function createCookie(cookieName, cookieValue) {
  let expirationDate = new Date();
  expirationDate.setTime(expirationDate.getTime() + (/*30 * 24 * 60*/ 20 * 60 * 1000)); // 20 minutes
  document.cookie = `${cookieName}=${cookieValue};expires=` + expirationDate.toUTCString();
}

function isJsonString(str) {
  try {
    JSON.parse(str);
  } catch (e) {
    return false;
  }
  return true;
}

function switchOverlay(connected) {
  if (connected) {
    document.getElementById("register_form").style.display = "none";
    document.getElementById("login_form").style.display = "none";
    document.getElementById("home_page").style.display = "grid";
    document.getElementById("create_post").style.display = "block";
    document.getElementById("display_posts").style.display = "block";


  } else {
    document.getElementById("register_form").style.display = "none";
    document.getElementById("login_form").style.display = "block";
    document.getElementById("home_page").style.display = "none";
    document.getElementById("create_post").style.display = "none";
    document.getElementById("display_posts").style.display = "none";
  }
}

$('register_form').show();
$('login_form').show();
