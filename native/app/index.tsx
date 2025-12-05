import { Text, View, StyleSheet, FlatList, ListRenderItem } from "react-native";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { BookstoreService, Book } from "@/gen/bookstore/v1/bookstore_pb";
import { useEffect, useState } from "react";

const HomeScreen = () => {
    const host = process.env.EXPO_PUBLIC_HOST;
    const port = process.env.EXPO_PUBLIC_PORT;

    const [books, setBooks] = useState<Book[]>([]);
    const [refreshing, setRefreshing] = useState(false);

    const client = createClient(BookstoreService, createConnectTransport({
        baseUrl: `http://${host}:${port}`,
    }));

    const fetchBooks = async () => {
        try {
            const response = await client.listBooks({
                shelf: "1", // Dummyなのでこの値で何か変わるということはない。
            });
            setBooks(response.books);
        } catch (err) {
            console.error("Failed to fetch books", err);
        }
    };

    const onRefresh = async () => {
        setRefreshing(true);
        await fetchBooks();
        setRefreshing(false);
    };

    useEffect(() => {
        fetchBooks();
    }, []);

    const renderItem: ListRenderItem<Book> = ({ item }) => (
        <View style={styles.itemContainer}>
            <Text style={styles.title}>書籍名：{item.name}</Text>
            <Text style={styles.author}>著者： {item.author}</Text>
        </View>
    );

    return (
        <View style={styles.container}>
            <FlatList
                data={books}
                renderItem={renderItem}
                keyExtractor={(item) => item.id}
                contentContainerStyle={styles.listContent}
                refreshing={refreshing}
                onRefresh={onRefresh}
            />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
    },
    listContent: {
        padding: 16,
    },
    itemContainer: {
        padding: 16,
        borderBottomWidth: 1,
        borderBottomColor: '#eee',
    },
    title: {
        fontSize: 18,
        fontWeight: 'bold',
    },
    author: {
        fontSize: 14,
        color: '#666',
        marginTop: 4,
    },
});

export default HomeScreen;
