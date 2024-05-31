import React from 'react';
import { View, Text, StyleSheet, Image } from 'react-native';
import { PieChart } from 'react-native-svg-charts';

const StatisticsScreen = () => {
    const data = [65, 25, 10];

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Доброе утро!</Text>
            <Text style={styles.subtitle}>Ваша статистика:</Text>
            <PieChart style={styles.chart} data={data} />
            <View style={styles.legend}>
                <View style={styles.legendItem}>
                    <View style={[styles.legendColor, { backgroundColor: '#6495ED' }]} />
                    <Text>Важные</Text>
                </View>
                <View style={styles.legendItem}>
                    <View style={[styles.legendColor, { backgroundColor: '#FF6347' }]} />
                    <Text>Здоровье</Text>
                </View>
                <View style={styles.legendItem}>
                    <View style={[styles.legendColor, { backgroundColor: '#D3D3D3' }]} />
                    <Text>Необязательные</Text>
                </View>
            </View>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 20,
        backgroundColor: '#fff',
    },
    title: {
        fontSize: 24,
        fontWeight: 'bold',
    },
    subtitle: {
        fontSize: 18,
        marginVertical: 10,
    },
    chart: {
        height: 200,
    },
    legend: {
        marginTop: 20,
    },
    legendItem: {
        flexDirection: 'row',
        alignItems: 'center',
        marginBottom: 10,
    },
    legendColor: {
        width: 20,
        height: 20,
        marginRight: 10,
    },
});

export default StatisticsScreen;