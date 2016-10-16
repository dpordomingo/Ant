<!DOCTYPE html>
<html>
    <head>
        <title>Ant Rover</title>
        <script>
            var source;
            var canvas;

            var start = function() {
                source = new EventSource('/events');
                source.onmessage = function(e) {
                    var data = e.data
                    draw(data)
                    console.log(data)
                };
            };

            var end = function() {
                source.close && source.close();
            };

            var draw = function(where) {
                coords = where.split(" ")
                ctx.lineTo(coords[0], coords[1]);
                ctx.stroke();
                ctx.beginPath();
                ctx.moveTo(coords[0], coords[1]);
            }

            window.onload = function() {
                var canvas = document.getElementById("Map");
                ctx = canvas.getContext("2d");
                ctx.beginPath();
                ctx.moveTo(0,0);
            }
        </script>
    </head>
    <body style="text-align:center;">
        <canvas id="Map" width="{{.Width}}" height="{{.Height}}" style="border:1px solid #000000;"></canvas>
        <br />
        <button onclick="javascript:start();">start</button>
        <button onclick="javascript:end();">end</button>
    </body>
</html>