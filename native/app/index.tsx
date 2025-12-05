import { Text, View, StyleSheet } from "react-native";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { BookstoreService } from "@/gen/bookstore/v1/bookstore_pb";
import { useEffect } from "react";

const HomeScreen = () => {
    const client = createClient(BookstoreService, createConnectTransport({
        baseUrl: "http://localhost:8080",
    }));

    useEffect(() => {
        client.listBooks({
            shelf: "shelf1",
        }).then((response) => {
            console.log(response);
        });
    }, []);

    return (
        <View style={styles.container}>
            <Text>Home</Text>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
    },
});

export default HomeScreen;
