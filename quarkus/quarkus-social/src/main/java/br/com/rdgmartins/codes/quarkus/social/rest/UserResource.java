package br.com.rdgmartins.codes.quarkus.social.rest;

import br.com.rdgmartins.codes.quarkus.social.dto.UserCreateRequest;
import br.com.rdgmartins.codes.quarkus.social.model.User;
import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;

@Path("/users")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
public class UserResource {

    @POST
    public Response userCreate(UserCreateRequest userCreateRequest) {

        User user = new User();
        user.setName(userCreateRequest.getName());
        user.setEmail(userCreateRequest.getEmail());
        user.setBirthday(userCreateRequest.getBirthdayDate());
        user.persist();

        return Response.ok(user).build();
    }

    @GET
    public Response listUsers() {
        return Response.ok().build();
    }

}
