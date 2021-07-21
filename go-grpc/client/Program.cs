using System;
using System.Threading.Tasks;
using System.Net.Http;

namespace client
{
    class Program
    {
        static async Task Main(String[]args){
            using (HttpClient client = new HttpClient()){
                var json = System.Text.Json.JsonSerializer.Serialize(new {
                    name = "Renato"
                });

                Console.WriteLine($"request data {json}");

                var content = new StringContent(json);

                content.Headers.ContentType = new System.Net.Http.Headers.MediaTypeHeaderValue("application/json");

                var resp = await client.PostAsync("http://localhost:8090/v1/example/echo",content);

                var res = await resp.Content.ReadAsStringAsync();

                Console.WriteLine($"response result {res}");
            }

            Console.ReadLine();
        }
    }
}
