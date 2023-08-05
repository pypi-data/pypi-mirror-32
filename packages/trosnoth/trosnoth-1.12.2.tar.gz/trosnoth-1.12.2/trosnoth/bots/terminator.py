from trosnoth.bots.goalsetter import Goal
from trosnoth.bots.ranger import RangerBot, HuntEnemies


class KillThings(Goal):
    '''
    Just kill things. Forever.
    '''
    def start(self):
        self.bot.player.onRemovedFromGame.addListener(self.removedFromGame)
        super(KillThings, self).start()

    def stop(self):
        super(KillThings, self).stop()
        self.bot.player.onRemovedFromGame.removeListener(self.removedFromGame)

    def removedFromGame(self, playerId):
        self.returnToParent()

    def reevaluate(self):
        self.bot.setUpgradePolicy(None)
        self.setSubGoal(HuntEnemies(self.bot, self))


class TerminatorBot(RangerBot):
    nick = 'TerminatorBot'
    generic = False

    MainGoalClass = KillThings
    pauseBetweenActions = False
    maxShotDelay = 0


BotClass = TerminatorBot

