MARGIN_OF_ERROR = 0.05
CONFIDENCE = 0.95
POPULATION = 1000
DISTRIBUTION = 0.5


class SampleSize(object):
    z_score = {0.80: 1.28,
               0.85: 1.44,
               0.90: 1.65,
               0.95: 1.96,
               0.99: 2.56}

    def __init__(self, population=None, confidence=None, margin=None):
        self.margin = MARGIN_OF_ERROR if not margin else margin
        self.confidence = CONFIDENCE if not confidence else confidence
        self.population = POPULATION if not population else population
        self.z_score_index = self.z_score[CONFIDENCE] if CONFIDENCE in self.z_score else None

    def sample_size(self):
        if not self.z_score_index:
            return "Usage of confidence should be [80, 85, 90, 95, 99]"
        return DISTRIBUTION * (1 - DISTRIBUTION) / ((self.margin / self.z_score_index) ** 2)

    def true_sample(self):
        ss = self.sample_size()
        return (ss * self.population) / (ss + self.population - 1)


if __name__ == '__main__':
    ss = SampleSize(population=1000000)
    print(ss.true_sample())
