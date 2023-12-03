using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using WEBSERVICE.Adapters;

namespace WEBSERVICE.Controllers;

[ApiController]
[Route("[controller]")]
public sealed class PokemonController(HttpClient client) : ControllerBase
{
    private readonly HttpClient _client = client;

    [HttpGet("{name}")]
    public async Task<IActionResult> GetAbilities(string name)
    {
        try
        {
            var response = await _client.GetAsync(
                $"https://ex.traction.one/pokedex/pokemon/{name}"
            );
            response.EnsureSuccessStatusCode();
            var body = JsonConvert.DeserializeObject<List<PokemonResponseAdapter>>(
                await response.Content.ReadAsStringAsync()
            );
            return Ok(body!);
        }
        catch (Exception e)
        {
            return NotFound(e.Message);
        }
    }
}
