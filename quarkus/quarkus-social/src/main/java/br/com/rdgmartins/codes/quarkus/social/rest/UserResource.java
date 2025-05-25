package br.com.rdgmartins.codes.quarkus.social.rest;

import br.com.rdgmartins.codes.quarkus.social.dto.UserCreateRequest;
import jakarta.ws.rs.Consumes;
import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;

@Path("/users")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
public class UserResource {

    @POST
    public Response userCreate(UserCreateRequest userCreateRequest) {
        return Response.ok().build();
    }

}
