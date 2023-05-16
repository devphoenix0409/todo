import { HttpLink, InMemoryCache, ApolloClient } from '@apollo/client/core';

export const apolloQueryClient = () => {
	const link = new HttpLink({
		uri: `http://localhost:8081/query`
	});

	const cache = new InMemoryCache();

	const apolloClient = new ApolloClient({
		link,
		cache
	});

	return apolloClient;
};
