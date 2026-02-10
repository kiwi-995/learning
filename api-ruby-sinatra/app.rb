# Minimal Todo API — Ruby / Sinatra + Sequel
# =============================================
# Run:  bundle install && ruby app.rb
# Test: curl http://localhost:4567/todos
#
# Key concepts demonstrated:
#   • Routing          — get/post/put/delete blocks
#   • JSON parsing     — JSON.parse(request.body.read)
#   • JSON responses   — .to_json on Ruby hashes/arrays
#   • ORM              — Sequel maps Ruby classes ↔ DB tables
#   • Content-Type     — telling the client "this is JSON"

require "sinatra"
require "sequel"
require "json"

# ─── Database setup ───────────────────────────────────────────────
# Sequel.connect returns a database object; SQLite stores in a file.
DB = Sequel.connect("sqlite://todos.db")

# Create the table if it doesn't exist.
# In Sequel, migrations look like this:
DB.create_table? :todos do
  primary_key :id
  String :title, null: false
  TrueClass :done, default: false     # Boolean column
end

# ─── ORM Model ────────────────────────────────────────────────────
# Sequel::Model maps the :todos table to a Ruby class.
# You get .create, .all, .first, .update, .destroy for free.
class Todo < Sequel::Model(:todos)
end

# ─── Helpers ──────────────────────────────────────────────────────
# Sinatra doesn't auto-parse JSON bodies, so we do it manually.
# This is a good thing to see — it shows you what frameworks hide!

def json_body
  JSON.parse(request.body.read, symbolize_names: true)
rescue JSON::ParserError
  halt 400, { error: "Invalid JSON" }.to_json
end

# Every response should be JSON
before do
  content_type :json
end

# ─── Routes ───────────────────────────────────────────────────────

# GET /todos — list all todos
get "/todos" do
  Todo.all.map { |t| serialize(t) }.to_json
end

# GET /todos/:id — get one todo
get "/todos/:id" do
  todo = Todo[params[:id]]
  halt 404, { error: "Todo not found" }.to_json unless todo
  serialize(todo).to_json
end

# POST /todos — create a new todo
post "/todos" do
  data = json_body
  halt 422, { error: "title is required" }.to_json unless data[:title]

  todo = Todo.create(title: data[:title], done: data[:done] || false)
  status 201
  serialize(todo).to_json
end

# PUT /todos/:id — update a todo
put "/todos/:id" do
  todo = Todo[params[:id]]
  halt 404, { error: "Todo not found" }.to_json unless todo

  data = json_body
  todo.update(
    title: data[:title] || todo.title,
    done:  data.key?(:done) ? data[:done] : todo.done
  )
  serialize(todo).to_json
end

# DELETE /todos/:id — delete a todo
delete "/todos/:id" do
  todo = Todo[params[:id]]
  halt 404, { error: "Todo not found" }.to_json unless todo

  todo.destroy
  status 204
  body ""
end

# ─── Serialisation helper ─────────────────────────────────────────
# Converts a Sequel model instance into a plain Ruby hash.
def serialize(todo)
  { id: todo.id, title: todo.title, done: todo.done }
end
