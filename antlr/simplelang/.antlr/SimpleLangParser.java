// Generated from /workspaces/Why.CSS/antlr/simplelang/SimpleLang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class SimpleLangParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, ID=12, INT=13, IMPORTNAMES=14, WS=15;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_expression = 2;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "expression"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'@use'", "'\"'", "';'", "'='", "'print'", "'('", "')'", "'*'", 
			"'/'", "'+'", "'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			"ID", "INT", "IMPORTNAMES", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SimpleLang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SimpleLangParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(7); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(6);
				statement();
				}
				}
				setState(9); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 4130L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	 
		public StatementContext() { }
		public void copyFrom(StatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentContext extends StatementContext {
		public TerminalNode ID() { return getToken(SimpleLangParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AssignmentContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintContext extends StatementContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PrintContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ImportContext extends StatementContext {
		public TerminalNode IMPORTNAMES() { return getToken(SimpleLangParser.IMPORTNAMES, 0); }
		public ImportContext(StatementContext ctx) { copyFrom(ctx); }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(27);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				_localctx = new ImportContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(11);
				match(T__0);
				setState(12);
				match(T__1);
				setState(13);
				match(IMPORTNAMES);
				setState(14);
				match(T__1);
				setState(15);
				match(T__2);
				}
				break;
			case ID:
				_localctx = new AssignmentContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(16);
				match(ID);
				setState(17);
				match(T__3);
				setState(18);
				expression(0);
				setState(19);
				match(T__2);
				}
				break;
			case T__4:
				_localctx = new PrintContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(21);
				match(T__4);
				setState(22);
				match(T__5);
				setState(23);
				expression(0);
				setState(24);
				match(T__6);
				setState(25);
				match(T__2);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	 
		public ExpressionContext() { }
		public void copyFrom(ExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MulDivContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public MulDivContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AddSubContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public AddSubContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParensContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ParensContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdContext extends ExpressionContext {
		public TerminalNode ID() { return getToken(SimpleLangParser.ID, 0); }
		public IdContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntContext extends ExpressionContext {
		public TerminalNode INT() { return getToken(SimpleLangParser.INT, 0); }
		public IntContext(ExpressionContext ctx) { copyFrom(ctx); }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(36);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				{
				_localctx = new IntContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(30);
				match(INT);
				}
				break;
			case ID:
				{
				_localctx = new IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(31);
				match(ID);
				}
				break;
			case T__5:
				{
				_localctx = new ParensContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(32);
				match(T__5);
				setState(33);
				expression(0);
				setState(34);
				match(T__6);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(46);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(44);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						_localctx = new MulDivContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(38);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(39);
						((MulDivContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__7 || _la==T__8) ) {
							((MulDivContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(40);
						expression(6);
						}
						break;
					case 2:
						{
						_localctx = new AddSubContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(41);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(42);
						((AddSubContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__9 || _la==T__10) ) {
							((AddSubContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(43);
						expression(5);
						}
						break;
					}
					} 
				}
				setState(48);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 2:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 5);
		case 1:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u000f2\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0001\u0000\u0004\u0000\b\b\u0000\u000b\u0000\f\u0000"+
		"\t\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001\u001c\b\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0003\u0002%\b\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002-\b\u0002\n\u0002\f\u0002"+
		"0\t\u0002\u0001\u0002\u0000\u0001\u0004\u0003\u0000\u0002\u0004\u0000"+
		"\u0002\u0001\u0000\b\t\u0001\u0000\n\u000b5\u0000\u0007\u0001\u0000\u0000"+
		"\u0000\u0002\u001b\u0001\u0000\u0000\u0000\u0004$\u0001\u0000\u0000\u0000"+
		"\u0006\b\u0003\u0002\u0001\u0000\u0007\u0006\u0001\u0000\u0000\u0000\b"+
		"\t\u0001\u0000\u0000\u0000\t\u0007\u0001\u0000\u0000\u0000\t\n\u0001\u0000"+
		"\u0000\u0000\n\u0001\u0001\u0000\u0000\u0000\u000b\f\u0005\u0001\u0000"+
		"\u0000\f\r\u0005\u0002\u0000\u0000\r\u000e\u0005\u000e\u0000\u0000\u000e"+
		"\u000f\u0005\u0002\u0000\u0000\u000f\u001c\u0005\u0003\u0000\u0000\u0010"+
		"\u0011\u0005\f\u0000\u0000\u0011\u0012\u0005\u0004\u0000\u0000\u0012\u0013"+
		"\u0003\u0004\u0002\u0000\u0013\u0014\u0005\u0003\u0000\u0000\u0014\u001c"+
		"\u0001\u0000\u0000\u0000\u0015\u0016\u0005\u0005\u0000\u0000\u0016\u0017"+
		"\u0005\u0006\u0000\u0000\u0017\u0018\u0003\u0004\u0002\u0000\u0018\u0019"+
		"\u0005\u0007\u0000\u0000\u0019\u001a\u0005\u0003\u0000\u0000\u001a\u001c"+
		"\u0001\u0000\u0000\u0000\u001b\u000b\u0001\u0000\u0000\u0000\u001b\u0010"+
		"\u0001\u0000\u0000\u0000\u001b\u0015\u0001\u0000\u0000\u0000\u001c\u0003"+
		"\u0001\u0000\u0000\u0000\u001d\u001e\u0006\u0002\uffff\uffff\u0000\u001e"+
		"%\u0005\r\u0000\u0000\u001f%\u0005\f\u0000\u0000 !\u0005\u0006\u0000\u0000"+
		"!\"\u0003\u0004\u0002\u0000\"#\u0005\u0007\u0000\u0000#%\u0001\u0000\u0000"+
		"\u0000$\u001d\u0001\u0000\u0000\u0000$\u001f\u0001\u0000\u0000\u0000$"+
		" \u0001\u0000\u0000\u0000%.\u0001\u0000\u0000\u0000&\'\n\u0005\u0000\u0000"+
		"\'(\u0007\u0000\u0000\u0000(-\u0003\u0004\u0002\u0006)*\n\u0004\u0000"+
		"\u0000*+\u0007\u0001\u0000\u0000+-\u0003\u0004\u0002\u0005,&\u0001\u0000"+
		"\u0000\u0000,)\u0001\u0000\u0000\u0000-0\u0001\u0000\u0000\u0000.,\u0001"+
		"\u0000\u0000\u0000./\u0001\u0000\u0000\u0000/\u0005\u0001\u0000\u0000"+
		"\u00000.\u0001\u0000\u0000\u0000\u0005\t\u001b$,.";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}