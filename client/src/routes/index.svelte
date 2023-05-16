<script context="module">
	export async function load({ stuff }) {
		try {
			const todos = await stuff.client.query({
				query: gql`
					query {
						todos {
							id: id
							text: text
						}
					}
				`
			});

			return {
				props: {
					todos: todos.data.todos
				}
			};
		} catch (err) {
			console.log('err: ', err);

			return {
				props: {
					todos: []
				}
			};
		}
	}
</script>

<script>
	import gql from 'graphql-tag';
	import { apolloQueryClient } from '$lib/apollo-client';
	import Modal from './modal.svelte';
	import { ModalType } from '../lib/model';

	let showModal = false;
	let action = ModalType.ADD;
	let modalControl;
	let currentID = -1;
	let currentText = '';

	export let todos;

	const handleAddTodo = () => {
		showModal = true;
		action = ModalType.ADD;
		currentText = '';
		currentID = -1;
	};

	const addTodo = async () => {
		try {
			const client = apolloQueryClient();
			const id = await client.mutate({
				mutation: gql`
					mutation {
						createTodo (
							todo:{
								text: "${currentText}",
								done: false
							}
						) {
							id: id
						}
					}
				`
			});

			todos = todos.concat({
				id: id.data.createTodo.id,
				text: currentText
			});
		} catch (err) {
			console.log('err: ', err);
		}

		currentText = '';
		modalControl.close();
	};

	const handleEditTodo = (/** @type {number} */ id) => {
		showModal = true;
		currentID = id;
		currentText = todos.filter((/** @type {{ id: number; }} */ todo) => todo.id === id)[0].text;
		action = ModalType.EDIT;
	};

	const updateTodo = async () => {
		try {
			const client = apolloQueryClient();
			await client.mutate({
				mutation: gql`
					mutation {
						updateTodo (
							id: ${currentID},
							changes: {
								text:"${currentText}"
							}
						) {
							id: id
						}
					}
				`
			});

			const updatedTodos = await client.query({
				query: gql`
					query {
						todos {
							id: id
							text: text
						}
					}
				`
			});

			todos = updatedTodos.data.todos;
		} catch (err) {
			console.log('err: ', err);
		}

		currentText = '';
		modalControl.close();
	};

	const closeModal = () => {
		modalControl.close();
		currentText = '';
		currentID = -1;
	};
</script>

<div class="page">
	<div class="todo-header">
		<h1>TODOS</h1>
		<button class="action-button" on:click={() => handleAddTodo()}>Add New Todo</button>
	</div>
	<section>
		<div>
			{#each todos as todo (todo.id)}
				<div class="todo-content">
					<div class="todo-data">
						<div class="todo-row">ID: {todo.id}</div>
						<div class="todo-row">Text: {todo.text}</div>
					</div>
					<div class="action">
						<button class="action-button" on:click={() => handleEditTodo(todo.id)}>Edit</button>
					</div>
				</div>
			{:else}
				<p>no todos...</p>
			{/each}
		</div>
	</section>

	<Modal bind:showModal bind:modalControl>
		<h2 slot="header">
			{action === ModalType.ADD ? 'Add a new Todo' : 'Edit Todo'}
		</h2>
		<div class="modal-content">
			<div class="p-10-0">Text:</div>
			<div>
				<input bind:value={currentText} class="action-button" type="text" />
			</div>
		</div>
		<hr />
		<div class="modal-footer">
			<button
				class="action-button"
				on:click={() => (action === ModalType.ADD ? addTodo() : updateTodo())}>Confirm</button
			>
			<button class="action-button" on:click={() => closeModal()}>Cancel</button>
		</div>
	</Modal>
</div>

<style>
	.page {
		padding: 1em;
		max-width: 768px;
		margin: 0 auto;
	}
	.todo-header {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
	.todo-content {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
	.todo-row {
		padding: 2.5px;
	}
	.todo-data {
		padding: 5px;
	}
	.action-button {
		min-width: 60px;
		height: 30px;
		cursor: pointer;
	}
	.modal-content {
		min-height: 9em;
	}
	.modal-footer {
		min-height: 3em;
		display: flex;
		justify-content: space-around;
		align-items: center;
	}
	.p-10-0 {
		padding: 10px 0 10px 0;
	}
</style>
