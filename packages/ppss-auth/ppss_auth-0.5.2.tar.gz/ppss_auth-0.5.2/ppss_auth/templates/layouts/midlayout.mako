<%inherit file="${context['supertpl']}" />

<div class="usermenu">
	<ul>
		<li><a href="${request.route_url('ppss:user:list')}">Users</a></li>
		<li><a href="${request.route_url('ppss:group:list')}">Groups</a></li>
		<li><a href="${request.route_url('ppss:perm:list')}">Permissions</a></li>
	</ul>

</div>
<div>


${next.body()}
</div>