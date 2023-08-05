<%inherit file="${context['midtpl']}" />


<form action="${request.route_url('ppss:group:edit',elementid=groupid)}" method="POST">
    
    <input type="text" name="name" placeholder="group name" value="${group.name if group else ""}">
    <br/>

    <label for="enablecheck">Enable:</label>
    <input id="enablecheck" type="checkbox" value="1" ${'checked="checked"' if group.enabled else ""}>
    <br/>
    <input type="submit" name="submit" value="Apply"/>

    <p>${msg}</p>
</form>

<div class="alluser">
    <ul>
    %for user in group.users:
            <li data-userid="${user.id}"> ${user.username}<span class="deleteuser">x</span></a></li>
    %endfor
    </ul>
</div>

<script type="text/javascript">
    $(".deleteuser").on("click",function(ev){
        $.ajax("${request.route_url('ppss:user:remove',userid=-1,groupid=group.id)}".replace("-1",$(this).attr("data-userid")  ),
            {datatype:"json",
            success:function(res){console.log("success");},
            error:function(res){console.log("error");}
            });
    });
</script>