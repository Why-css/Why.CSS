// Generated from /workspaces/java-env/simplelang/SimpleLang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SimpleLangParser}.
 */
public interface SimpleLangListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SimpleLangParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(SimpleLangParser.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleLangParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(SimpleLangParser.ProgContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleLangParser#stat}.
	 * @param ctx the parse tree
	 */
	void enterStat(SimpleLangParser.StatContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleLangParser#stat}.
	 * @param ctx the parse tree
	 */
	void exitStat(SimpleLangParser.StatContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterMulDiv(SimpleLangParser.MulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitMulDiv(SimpleLangParser.MulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(SimpleLangParser.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(SimpleLangParser.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterParens(SimpleLangParser.ParensContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitParens(SimpleLangParser.ParensContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Id}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterId(SimpleLangParser.IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Id}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitId(SimpleLangParser.IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Int}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterInt(SimpleLangParser.IntContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Int}
	 * labeled alternative in {@link SimpleLangParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitInt(SimpleLangParser.IntContext ctx);
}