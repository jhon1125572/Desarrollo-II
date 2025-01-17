{{define "login"}}
<div class="modal fade" id="LoginModal" tabindex="-1" role="dialog" aria-labelledby="LoginModalTitle" aria-hidden="true">
    <div class="modal-dialog modal-sm">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title" id="LoginModalTitle">INICIAR SESIÓN</h4>
            </div>
            <div class="modal-body">
                <form class="login-form" id="LoginForm" name="LoginForm" onkeypress="return checkEnter(event)">
                    <div class="form-group ">
                        <input type="text" id="Username" name="Username" class="form-control" placeholder="username">
                        <i class="glyphicon glyphicon-user"></i>
                    </div>
                    <div class="form-group">
                        <input type="password" id="Password" name="Password" class="form-control" placeholder="Contraseña">
                        <i class="glyphicon glyphicon-lock"></i>
                    </div>
                </form>
            </div>
            <div class="modal-footer">
                <button type="button" id="LoginButton" name="LoginButton" class="btn btn-primary" data-dismiss="modal" aria-label="Left Align">
                    <i class="glyphicon glyphicon-ok"></i>
                    Iniciar
                </button>
                <button type="button" class="btn btn-default" data-dismiss="modal">Cerrar</button>
            </div>
        </div>
    </div>
</div>
{{ end }}
