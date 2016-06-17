{{define "header"}}
<!DOCTYPE html>
<html lang="es">
    <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width; initial-scale=1.0"> 
    <title>{{.Title}} - PersonalCollectionMovies [WSPCM]</title>
        <!-- Estilos CSS vinculados -->
         <link href="//getbootstrap.com/dist/css/bootstrap.css" rel="stylesheet">
        <link href="//netdna.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" rel='stylesheet' type='text/css'>
        <!--
        <link href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.2/css/bootstrap.min.css" rel="stylesheet">
        <link href="//netdna.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" rel='stylesheet' type='text/css'>
        <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css">
        -->
        <link href="//maxcdn.bootstrapcdn.com/font-awesome/4.2.0/css/font-awesome.min.css" rel="stylesheet">
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.5.0/css/bootstrap-datepicker.min.css" />
        <link href="//oss.maxcdn.com/jquery.bootstrapvalidator/0.5.2/css/bootstrapValidator.min.css" rel="stylesheet"/>
        <link href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-rating/1.3.2/bootstrap-rating.css" rel="stylesheet"/>
        <link href="https://cdn.datatables.net/1.10.12/css/dataTables.bootstrap.min.css" rel="stylesheet"/>
        <link href="//cdn.datatables.net/1.10.12/css/jquery.dataTables.min.css" rel="stylesheet"/>
        <!--
        <link href="//cdnjs.cloudflare.com/ajax/libs/formvalidation/0.6.1/css/formValidation.min.css" rel="stylesheet"/>
        <link href="//oss.maxcdn.com/jquery.bootstrapvalidator/0.5.2/css/bootstrapValidator.min.css" rel="stylesheet"/>
        -->
        <!-- Estilos propios -->
        <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/footer.css" rel="stylesheet">
        <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/footer-basic-centered.css" rel="stylesheet">
        <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/header.css" rel="stylesheet">
        <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/login.css" rel="stylesheet">
        <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/general.css" rel="stylesheet">
        <link href="//personalcollectionmovies-alobaton.c9users.io/public/css/dropdown.css" rel="stylesheet">
    </head>
    <body onload="checkCookies()">
        <section id="preloader">
    		<div class="loading-circle fa-spin"></div>
    	</section>
        <!-- Barra de navegación estatica-->
        <!-- Debe contener inicio de seión y registro -->
        <nav class="navbar navbar-fixed-top navbar-inverse">
            <div class="container">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#NavbarCollapse" aria-expanded="true">
                        <i class="glyphicon glyphicon-menu-hamburger"></i>
                    </button>
                    <a href="//personalcollectionmovies-alobaton.c9users.io/" class="navbar-brand"><strong>PersonalCollectionMovies</strong></a>
                </div>
                <div id="NavbarCollapse" name="NavbarCollapse" class="collapse navbar-collapse">
                    <!-- Definimos una variable que nos permita cambiar los controles de inicio y cierre de sesión -->
                    {{.SessionControl}}
                </div>
            </div>
        </nav>
        {{template "login" .}}
        <!-- Fin barra de navegacion -->
        <!-- Carrusel -->
        <div class="container">
            <div id="Carousel" class="carousel slide" data-ride="carousel">
                <!-- Indicators -->
                <ol class="carousel-indicators">
                    <li data-target="#Carousel" data-slide-to="0" class="active"></li>
                    <li data-target="#Carousel" data-slide-to="1"></li>
                </ol>
                <!-- Wrapper for slides -->
                <div class="carousel-inner" role="listbox">
                    <div class="item active">
                        <img src="//personalcollectionmovies-alobaton.c9users.io/public/images/slide%20(1).jpg">
                        <div class="carousel-caption">
                        </div>
                    </div>
                    <div class="item">
                        <img src="//personalcollectionmovies-alobaton.c9users.io/public/images/slide%20(2).jpg">
                        <div class="carousel-caption">
                        </div>
                    </div>
                </div>
                <!-- Controls -->
                <a class="left carousel-control" href="#Carousel" role="button" data-slide="prev">
                    <span class="glyphicon glyphicon-chevron-left" aria-hidden="true"></span>
                    <span class="sr-only">Previous</span>
                </a>
                <a class="right carousel-control" href="#Carousel" role="button" data-slide="next">
                    <span class="glyphicon glyphicon-chevron-right" aria-hidden="true"></span>
                    <span class="sr-only">Next</span>
                </a>
            </div>
        </div>
        <!-- Carrusel -->
        <!-- Barra de navegación de pagínas -->
        <div class="container">
            <nav class="navbar navbar-inverse navbar">
                <div class="container">
                    <div class="navbar-header">
                        <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navbarMainMenu" aria-expanded="true">
                            <i class="glyphicon glyphicon-menu-hamburger"></i>
                        </button>
                    </div>
                    <div id="navbarMainMenu"  class="collapse navbar-collapse">
                        <ul class="nav navbar-nav">
                            <li><a href="//personalcollectionmovies-alobaton.c9users.io/">INICIO</a></li>
                            <li><a href="#">RECIENTES</a></li>
                            <li><a href="#">CARTELERA</a></li> 
                            <li><a href="#">ESTRENOS</a></li> 
                            <li><a href="#">POPULARES</a></li> 
                            <li><a href="#">CONOCENOS</a></li> 
                            <li><a href="#">CONTACTENOS</a></li> 
                        </ul>
                    </div>
                </div>
            </nav>
        </div>
        <!-- Fin barra de navegacion paginas -->
        <!-- Buscar pelicula -->
        <div class="container">
            <div class="input-group">
                <span class="input-group-btn">
                   <button class="btn btn-primary" type="button">BUSCAR PELICULA</button>
                </span>
                <input type="text" class="form-control" placeholder="Nombre de la pelicula">
            </div>
            <div class="container">
            </div>
        </div>
    <!-- Fin header -->
{{end}}