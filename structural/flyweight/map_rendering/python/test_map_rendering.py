import unittest
from map_rendering import (
    MapRenderer,
    VehicleType,
    VehicleStatus,
    VehicleIconFactory
)


class TestMapRendering(unittest.TestCase):
    def setUp(self):
        self.renderer = MapRenderer()
        # Clear the icon factory cache
        VehicleIconFactory._icons.clear()

    def test_add_vehicle(self):
        self.renderer.add_vehicle("car1", VehicleType.CAR)
        self.assertEqual(self.renderer.get_vehicle_count(), 1)

    def test_update_vehicle(self):
        self.renderer.add_vehicle("car1", VehicleType.CAR)
        self.renderer.update_vehicle(
            "car1",
            position=(37.7749, -122.4194),
            speed=45.0,
            status=VehicleStatus.MOVING,
            heading=90.0
        )
        render_output = self.renderer.render_map()[0]
        self.assertIn("car1", render_output)
        self.assertIn("(37.7749, -122.4194)", render_output)

    def test_icon_sharing(self):
        # Add multiple vehicles of the same type
        self.renderer.add_vehicle("car1", VehicleType.CAR)
        self.renderer.add_vehicle("car2", VehicleType.CAR)
        self.renderer.add_vehicle("bus1", VehicleType.BUS)

        # Check that we have 3 vehicles but only 2 unique icons
        self.assertEqual(self.renderer.get_vehicle_count(), 3)
        self.assertEqual(self.renderer.get_icon_count(), 2)

    def test_render_map(self):
        self.renderer.add_vehicle("car1", VehicleType.CAR)
        self.renderer.add_vehicle("bus1", VehicleType.BUS)
        self.renderer.update_vehicle(
            "car1",
            position=(37.7749, -122.4194),
            speed=45.0,
            status=VehicleStatus.MOVING,
            heading=90.0
        )
        self.renderer.update_vehicle(
            "bus1",
            position=(37.7749, -122.4194),
            speed=30.0,
            status=VehicleStatus.STOPPED,
            heading=180.0
        )

        render_outputs = self.renderer.render_map()
        self.assertEqual(len(render_outputs), 2)
        self.assertTrue(any("car1" in output for output in render_outputs))
        self.assertTrue(any("bus1" in output for output in render_outputs))


if __name__ == '__main__':
    unittest.main() 