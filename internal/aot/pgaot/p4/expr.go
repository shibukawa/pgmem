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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v6 = int32(4549024)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v9
	v12 = F_palloc0(m, int32(72))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v12))) = int64(382)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v18
		v24 = F_AllocSetContextCreateInternal(m, v18, int32(67362), int32(0), int32(8192), int32(8388608))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v24
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v27
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			v30 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v30
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l0
			v33 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+52)) = uint8(v33)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v30
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)) = uint8(v33)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v30
			*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v29
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			v45 = F_lcons(m, v12, v44)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v45
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v7
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == v3 {
		v81 = v3
		m.G0 = v7 + int32(16)
		return v81
	} else {
		v12 = F_palloc0(m, int32(68))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(380)
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
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
						if v36 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(16)
							v42 = F_palloc(m, int32(640))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v55 = v42
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
								v57 = v55
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v58 + int32(1)
								v64 = v57 + v58*int32(40)
								v65 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v64))) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+24)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+16)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+8)) = v65
								v75 = F_jit_compile_expr(m, v12)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									if v75 != 0 {
										v81 = v12
										m.G0 = v7 + int32(16)
										return v81
									} else {
										F_ExecReadyInterpretedExpr(m, v12)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v81 = v12
											m.G0 = v7 + int32(16)
											return v81
										}
									}
								}
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
							if v44 != v36 {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
								v57 = v46
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v58 + int32(1)
								v64 = v57 + v58*int32(40)
								v65 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v64))) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+24)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+16)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+8)) = v65
								v75 = F_jit_compile_expr(m, v12)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									if v75 != 0 {
										v81 = v12
										m.G0 = v7 + int32(16)
										return v81
									} else {
										F_ExecReadyInterpretedExpr(m, v12)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v81 = v12
											m.G0 = v7 + int32(16)
											return v81
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v36 << (uint(int32(1)) % 32)
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
								v53 = F_repalloc(m, v50, v36*int32(80))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = v53
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
									v57 = v55
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v58 + int32(1)
									v64 = v57 + v58*int32(40)
									v65 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v64))) = v65
									*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v65
									*(*int64)(unsafe.Add(mBase, uint32(v64)+24)) = v65
									*(*int64)(unsafe.Add(mBase, uint32(v64)+16)) = v65
									*(*int64)(unsafe.Add(mBase, uint32(v64)+8)) = v65
									v75 = F_jit_compile_expr(m, v12)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										if v75 != 0 {
											v81 = v12
											m.G0 = v7 + int32(16)
											return v81
										} else {
											F_ExecReadyInterpretedExpr(m, v12)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												v81 = v12
												m.G0 = v7 + int32(16)
												return v81
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
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v3 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v29
L2:
	;
	v10 = v3
	v11 = v3
	goto L7
L3:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v5 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v29 = v3
	goto L1
L6:
	;
	goto L5
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v11<<(uint(int32(2))%32))))
	v17 = F_ExecInitExpr(m, v16, l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v29 = v21
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v21 = F_lappend(m, v10, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v24 = v11 + int32(1)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v24 < v25 {
		v10 = v21
		v11 = v24
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == v3 {
		v81 = v3
		m.G0 = v7 + int32(16)
		return v81
	} else {
		v12 = F_palloc0(m, int32(68))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(380)
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
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
						if v36 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(16)
							v42 = F_palloc(m, int32(640))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v55 = v42
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
								v57 = v55
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v58 + int32(1)
								v64 = v57 + v58*int32(40)
								v65 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v64))) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+24)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+16)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+8)) = v65
								v75 = F_jit_compile_expr(m, v12)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									if v75 != 0 {
										v81 = v12
										m.G0 = v7 + int32(16)
										return v81
									} else {
										F_ExecReadyInterpretedExpr(m, v12)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v81 = v12
											m.G0 = v7 + int32(16)
											return v81
										}
									}
								}
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
							if v44 != v36 {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
								v57 = v46
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v58 + int32(1)
								v64 = v57 + v58*int32(40)
								v65 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v64))) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+24)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+16)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v64)+8)) = v65
								v75 = F_jit_compile_expr(m, v12)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									if v75 != 0 {
										v81 = v12
										m.G0 = v7 + int32(16)
										return v81
									} else {
										F_ExecReadyInterpretedExpr(m, v12)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v81 = v12
											m.G0 = v7 + int32(16)
											return v81
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v36 << (uint(int32(1)) % 32)
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
								v53 = F_repalloc(m, v50, v36*int32(80))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = v53
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
									v57 = v55
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v58 + int32(1)
									v64 = v57 + v58*int32(40)
									v65 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v64))) = v65
									*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v65
									*(*int64)(unsafe.Add(mBase, uint32(v64)+24)) = v65
									*(*int64)(unsafe.Add(mBase, uint32(v64)+16)) = v65
									*(*int64)(unsafe.Add(mBase, uint32(v64)+8)) = v65
									v75 = F_jit_compile_expr(m, v12)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										if v75 != 0 {
											v81 = v12
											m.G0 = v7 + int32(16)
											return v81
										} else {
											F_ExecReadyInterpretedExpr(m, v12)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												v81 = v12
												m.G0 = v7 + int32(16)
												return v81
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
func F_RegisterExprContextCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	v7 = F_MemoryContextAlloc(m, v5, int32(12))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v7
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
	var v47 int32
	_ = v47
	v3 = int32(0)
	if l0 == v3 {
		v47 = v3
		return v47
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = v7 - int32(11)
		if base.Ui32(int32(9)) < base.Ui32(v9) {
			v47 = v3
			return v47
		} else {
			if int32(base.Ui32(int32(977))>>(uint(v9)%32))&int32(1) == int32(0) {
				v47 = v3
				return v47
			} else {
				if l1 < int32(0) {
					v47 = v3
					return v47
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v9<<(uint(int32(2))%32))+uint32(_consts[1357])))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24)))
					if v26 == int32(0) {
						v47 = v3
						return v47
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						if v29 <= l1 {
							v47 = v3
							return v47
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
									v47 = v36
									return v47
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if v42 != int32(20) {
										v47 = v36
										return v47
									} else {
										v45 = F_get_base_element_type(m, v36)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											v47 = v45
											return v47
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
