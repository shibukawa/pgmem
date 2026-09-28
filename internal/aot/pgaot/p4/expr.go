package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAssignExprContext(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v6 = int32(_a_F_ExecAssignExprContext_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAssignExprContext[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecAssignExprContext[0])) = v9
	v12 = F_palloc0(m, int32(88))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v12))) = int64(388)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v18
		v24 = F_AllocSetContextCreateInternal(m, v18, int32(_a_F_ExecAssignExprContext_1), int32(0), int32(_a_F_ExecAssignExprContext_2), int32(_a_F_ExecAssignExprContext_3))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v24
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v27
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			v30 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v30
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v29
			*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v30
			*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = l0
			v38 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+64)) = uint8(v38)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v30
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)) = uint8(v38)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			v45 = F_lcons(m, v12, v44)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v45
				*(*int32)(unsafe.Add(mBase, _c_F_ExecAssignExprContext[0])) = v7
				*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v12
				return
			}
		}
	}
}
func F_ExecInitExpr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == v3 {
		v82 = v3
		m.G0 = v7 + int32(16)
		return v82
	} else {
		v12 = F_palloc0(m, int32(72))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(386)
			v22 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v22
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v22
			v26 = F_expr_setup_walker(m, l0, v7)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_ExecPushExprSetupSteps(m, v12, v7)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_ExecInitExprRec(m, l0, v12, v12+int32(8), v12+int32(5))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
						if v36 == int32(0) {
							v39 = int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v39
							v43 = F_palloc_mul(m, int32(40), v39)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v56 = v43
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v56
								v58 = v56
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v59 + int32(1)
								v65 = v58 + v59*int32(40)
								v66 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
								v76 = F_jit_compile_expr(m, v12)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									if v76 != 0 {
										v82 = v12
										m.G0 = v7 + int32(16)
										return v82
									} else {
										F_ExecReadyInterpretedExpr(m, v12)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v82 = v12
											m.G0 = v7 + int32(16)
											return v82
										}
									}
								}
							}
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
							if v45 != v36 {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
								v58 = v47
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v59 + int32(1)
								v65 = v58 + v59*int32(40)
								v66 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
								v76 = F_jit_compile_expr(m, v12)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									if v76 != 0 {
										v82 = v12
										m.G0 = v7 + int32(16)
										return v82
									} else {
										F_ExecReadyInterpretedExpr(m, v12)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v82 = v12
											m.G0 = v7 + int32(16)
											return v82
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v36 << (uint(int32(1)) % 32)
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
								v54 = F_repalloc(m, v51, v36*int32(80))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v54
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v56
									v58 = v56
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v59 + int32(1)
									v65 = v58 + v59*int32(40)
									v66 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
									v76 = F_jit_compile_expr(m, v12)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										if v76 != 0 {
											v82 = v12
											m.G0 = v7 + int32(16)
											return v82
										} else {
											F_ExecReadyInterpretedExpr(m, v12)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												v82 = v12
												m.G0 = v7 + int32(16)
												return v82
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecInitExprList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v16 = v3
	v17 = v3
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v17<<(uint(int32(2))%32))))
	v23 = F_ExecInitExpr(m, v22, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return v27
L9:
	;
	return int32(0)
L10:
	;
	v27 = F_lappend(m, v16, v23)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v30 = v17 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 < v31 {
		v16 = v27
		v17 = v30
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
}
func F_ExecInitExprWithParams(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == v3 {
		v82 = v3
		m.G0 = v7 + int32(16)
		return v82
	} else {
		v12 = F_palloc0(m, int32(72))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(386)
			v22 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v22
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v22
			v26 = F_expr_setup_walker(m, l0, v7)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_ExecPushExprSetupSteps(m, v12, v7)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_ExecInitExprRec(m, l0, v12, v12+int32(8), v12+int32(5))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
						if v36 == int32(0) {
							v39 = int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v39
							v43 = F_palloc_mul(m, int32(40), v39)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v56 = v43
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v56
								v58 = v56
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v59 + int32(1)
								v65 = v58 + v59*int32(40)
								v66 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
								v76 = F_jit_compile_expr(m, v12)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									if v76 != 0 {
										v82 = v12
										m.G0 = v7 + int32(16)
										return v82
									} else {
										F_ExecReadyInterpretedExpr(m, v12)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v82 = v12
											m.G0 = v7 + int32(16)
											return v82
										}
									}
								}
							}
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
							if v45 != v36 {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
								v58 = v47
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v59 + int32(1)
								v65 = v58 + v59*int32(40)
								v66 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v66
								*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
								v76 = F_jit_compile_expr(m, v12)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									if v76 != 0 {
										v82 = v12
										m.G0 = v7 + int32(16)
										return v82
									} else {
										F_ExecReadyInterpretedExpr(m, v12)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v82 = v12
											m.G0 = v7 + int32(16)
											return v82
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v36 << (uint(int32(1)) % 32)
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
								v54 = F_repalloc(m, v51, v36*int32(80))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v54
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v56
									v58 = v56
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v59 + int32(1)
									v65 = v58 + v59*int32(40)
									v66 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v66
									*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
									v76 = F_jit_compile_expr(m, v12)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										if v76 != 0 {
											v82 = v12
											m.G0 = v7 + int32(16)
											return v82
										} else {
											F_ExecReadyInterpretedExpr(m, v12)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												v82 = v12
												m.G0 = v7 + int32(16)
												return v82
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_RegisterExprContextCallback(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = F_MemoryContextAlloc(m, v5, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
		*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v7
		return
	}
}
func F_get_call_expr_argtype(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v3 = int32(0)
	if l0 == v3 {
		v48 = v3
		return v48
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = v7 - int32(11)
		v16 = int32(0)
		if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v9))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v9)%32))&int32(1) == v16)|base.B2i32(l1 < v16) != 0 {
			v48 = v3
			return v48
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v9<<(uint(int32(2))%32))+uint32(_c_F_get_call_expr_argtype[0])))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24)))
			if v26 == int32(0) {
				v48 = v3
				return v48
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
				if v29 <= l1 {
					v48 = v3
					return v48
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+l1<<(uint(int32(2))%32))))
					v36 = F_exprType(m, v35)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						if l1 != int32(1) {
							v48 = v36
							return v48
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v42 != int32(20) {
								v48 = v36
								return v48
							} else {
								v45 = F_get_base_element_type(m, v36)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v48 = v45
									return v48
								}
							}
						}
					}
				}
			}
		}
	}
}
