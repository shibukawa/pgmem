package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_coerce_fn_result_column(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	if l3 == int32(0) {
		v32 = int32(0)
		v35 = F_makeVarFromTargetEntry(m, int32(1), l0)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
			v41 = F_coerce_to_target_type(m, v32, v35, v37, l1, l2, int32(1), int32(2), int32(-1))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				if v41 == int32(0) {
					v75 = v32
					return v75
				} else {
					F_assign_expr_collations(m, int32(0), v41)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						if v41 == v35 {
							v52 = v41
						} else {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v49)
							v52 = v41
						}
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						if v57 != 0 {
							v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
							v62 = v58 + int32(1)
						} else {
							v62 = int32(1)
						}
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v66 = F_makeTargetEntry(m, v52, base.I32_extend16_s(v62), v64, int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							v69 = F_lappend(m, v68, v66)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v69
								v75 = int32(1)
								return v75
							}
						}
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v10 != 0 {
			v32 = int32(0)
			v35 = F_makeVarFromTargetEntry(m, int32(1), l0)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
				v41 = F_coerce_to_target_type(m, v32, v35, v37, l1, l2, int32(1), int32(2), int32(-1))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					if v41 == int32(0) {
						v75 = v32
						return v75
					} else {
						F_assign_expr_collations(m, int32(0), v41)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v41 == v35 {
								v52 = v41
							} else {
								v49 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v49)
								v52 = v41
							}
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							if v57 != 0 {
								v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
								v62 = v58 + int32(1)
							} else {
								v62 = int32(1)
							}
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v66 = F_makeTargetEntry(m, v52, base.I32_extend16_s(v62), v64, int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								v69 = F_lappend(m, v68, v66)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v69
									v75 = int32(1)
									return v75
								}
							}
						}
					}
				}
			}
		} else {
			v11 = int32(0)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = F_exprType(m, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v21 = F_coerce_to_target_type(m, v11, v13, v14, l1, l2, int32(1), int32(2), int32(-1))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v21 == int32(0) {
						v75 = v11
						return v75
					} else {
						F_assign_expr_collations(m, int32(0), v21)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
							v30 = F_makeVarFromTargetEntry(m, int32(1), l0)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v52 = v30
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								if v57 != 0 {
									v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
									v62 = v58 + int32(1)
								} else {
									v62 = int32(1)
								}
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v66 = F_makeTargetEntry(m, v52, base.I32_extend16_s(v62), v64, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									v69 = F_lappend(m, v68, v66)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v69
										v75 = int32(1)
										return v75
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
func F_get_fn_expr_variadic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = int32(0)
	if l0 == v2 {
		v13 = v2
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v13 = v2
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int32(15) {
				v13 = v2
			} else {
				v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+13)))
				v13 = v11
			}
		}
	}
	return v13 & int32(1)
}
func Fn13822(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		F_errmsg_internal(m, l3, v8)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(_a_Fn13822_0), l2, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func Fn13831(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+6)))
	v12 = int32(2)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8*int32(0)<<(uint(v12)%32)+l1<<(uint(v12)%32)-int32(4))))
	if v20 == int32(0) {
		v33 = l2
		return v33
	} else {
		v24 = F_index_getprocinfo(m, l0, int32(1), l1)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v24 == int32(0) {
				v33 = l2
				return v33
			} else {
				v30 = F_FunctionCall0Coll(m, v24)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v33 = v30
					return v33
				}
			}
		}
	}
}
func Fn13842(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_Fn13842_0), int32(0))
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(_a_Fn13842_1), int32(0))
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13842_2), l2, l1)
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func Fn13848(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_DatumGetAnyArrayP(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = F_array_contain_compare(m, v7, v12, v14, l1, v15+int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				if v20 == int32(-1) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v27 == int32(-1) {
						return v18
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v12 == v30 {
							return v18
						} else {
							F_pfree(m, v12)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								return v18
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v7 == v23 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v27 == int32(-1) {
							return v18
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v12 == v30 {
								return v18
							} else {
								F_pfree(m, v12)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									return v18
								}
							}
						}
					} else {
						F_pfree(m, v7)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							if v27 == int32(-1) {
								return v18
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v12 == v30 {
									return v18
								} else {
									F_pfree(m, v12)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int32(0)
									} else {
										return v18
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
func Fn13851(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13851[0])))
	if v6 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_Fn13851_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13851_1), l3, l2)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
		return int32(0)
	}
}
func Fn13857(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v22 = v20 & int32(1)
			if v22 != 0 {
				v23 = v14
			} else {
				v23 = v9 + int32(4)
			}
			if v20 == int32(1) {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v29 == int32(18) {
					v32 = int32(16)
				} else {
					v32 = int32(0)
				}
				if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v39 = int32(4)
				} else {
					v39 = v32
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v22 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(1)
			v52 = v16 + v51
			v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v57 = v55 & v51
			if v57 != 0 {
				v58 = v52
			} else {
				v58 = v16 + int32(4)
			}
			if v55 == int32(1) {
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
				if v64 == int32(18) {
					v67 = int32(16)
				} else {
					v67 = int32(0)
				}
				if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v74 = int32(4)
				} else {
					v74 = v67
				}
				v85 = v74
			} else {
				v75 = int32(1)
				if v57 != 0 {
					v85 = int32(base.Ui32(v55)>>(uint(v75)%32)) - v75
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v86 = F_dotrim(m, v23, v50, v58, v85, l2, l1)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				return v86
			}
		}
	}
}
func Fn13866(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = F_strlen(m, l0)
	mBase = m.M
	if v18 != 0 {
		v20 = F_strstr(m, l0, int32(_a_Fn13866_0))
		mBase = m.M
		if v20 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l0
					F_errmsg(m, l4, v14+int32(-16))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						F_errdetail(m, l8, int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_Fn13866_1), l7, l1)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v21 == int32(45) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l0
						F_errmsg(m, l4, v14+int32(-48))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							F_errdetail(m, l6, int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_Fn13866_1), l5, l1)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v18-int32(1)))))
				if v27 == int32(45) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l0
							F_errmsg(m, l4, v14+int32(-48))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								F_errdetail(m, l6, int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_Fn13866_1), l5, l1)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					v31 = Fn13880(m, l0, int32(47))
					mBase = m.M
					if v31 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l0
								F_errmsg(m, l4, v14+int32(-32))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									F_errdetail(m, l3, int32(0))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_Fn13866_1), l2, l1)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					} else {
						m.G0 = v16 - int32(-64)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
				F_errmsg(m, l4, v16)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errdetail(m, l10, int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_Fn13866_1), l9, l1)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func Fn13873(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_get_fn_expr_argtype(m, v12, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_enum_endpoint(m, v14, l4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = F_format_type_be(m, v14)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v47
								F_errmsg(m, int32(_a_Fn13873_0), v10)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13873_1), l2, l1)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					m.G0 = v10 + int32(16)
					return v18
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_Fn13873_2), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_Fn13873_1), l3, l1)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func Fn13886(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = F_gbt_num_union(m, v7, v6, l1, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func Fn13893(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = F_DirectFunctionCall2Coll(m, l4, int32(0), v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v18 = v9
			return v18
		} else {
			v16 = F_DirectFunctionCall2Coll(m, l4, int32(0), v7+l3, v8+l3)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = v16
				return v18
			}
		}
	}
}
func Fn13895(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v22 = v11 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v25 = int32(4)
		v26 = v23 + v25
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v29 = int32(2)
		v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
			v42 = v26 + (v30+int32(3))&int32(2147483644)
		} else {
			v42 = v26
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v42
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v44)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v46 == v44 {
			v50 = *(*int32)(unsafe.Add(mBase, _c_Fn13895[0]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v51*int32(28))+uint32(_c_Fn13895[1])))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v56
		} else {
		}
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+16)))
		v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+v64)+12)))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v70 = F_gbt_var_consistent(m, v11+int32(8), v15, v19&int32(_a_Fn13895_0), v62, v66&int32(1), l1, v69)
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			m.G0 = v11 + int32(16)
			return v70
		}
	}
}
func Fn13901(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v9 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v17 = F_GetSysCacheOid(m, l7, l0, v9, v9, v9)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if l1|v17 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errcode(m, l6)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
					F_errmsg(m, l5, v12)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, l4, l3, l2)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			m.G0 = v12 + int32(16)
			return v17
		}
	}
}
func Fn13910(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v20 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = int32(0)
	goto L4
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v25<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = l0
	v43 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)) = uint16(v43)
	v46 = v25 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v46 < v47 {
		v25 = v46
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L5
L7:
	;
	m.G0 = v16 + int32(16)
	return
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v64 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v69 = int32(0)
	goto L10
L10:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v69<<(uint(int32(2))%32))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if base.Ui32(l10) < base.Ui32(v86) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L7
L12:
	;
	v116 = v69 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v116 < v117 {
		v69 = v116
		goto L10
	} else {
		goto L23
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = l0
	v113 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v85)+24)) = uint16(v113)
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v89 = int32(1) << (uint(v86) % 32)
	if v89&l9 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v89&l8 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+24)) = uint8(v94)
	goto L12
L18:
	;
	return
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v104
	F_errmsg(m, int32(_a_Fn13910_0), v16)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, l6, l5, l4)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	goto L11
}
func Fn13916(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v20)
		v22 = F_convert_any_priv_string(m, v16, l1)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, l2, v13, v14, v22, v11+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v11 + int32(16)
				return v35
			}
		}
	}
}
func Fn13923(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v14 = int32(0)
	goto L3
L3:
	;
	v15 = F_list_concat_copy(m, l2, l3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	v11 = F_get_opclass_input_type(m, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v14 = v11
	goto L3
L7:
	;
	return
L8:
	;
	if v15 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v24 = l1
	v25 = int32(0)
	v29 = v14
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v25<<(uint(int32(2))%32))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L7
L13:
	;
	v63 = v25 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v63 < v64 {
		v24 = v59
		v25 = v63
		v29 = v61
		goto L11
	} else {
		goto L27
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = l0
	v57 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v57)
	v59 = v24
	v61 = v29
	goto L13
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v37 != int32(1) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	if v40 != v41 {
		goto L14
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	if v40 != v29 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v44 = F_opclass_for_family_datatype(m, l4, l0, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v46 = v24
	v47 = v29
	goto L22
L22:
	;
	if v46 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v46 = v44
	v47 = v40
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v46
	v49 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v49)
	v59 = v46
	v61 = v47
	goto L13
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = l0
	v52 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v52)
	v59 = int32(0)
	v61 = v47
	goto L13
L27:
	;
	goto L12
}
func Fn13925(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v261 int64
	_ = v261
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int64
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v308 int64
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int64
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int64
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int64
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int64
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int64
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int64
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v388 int64
	_ = v388
	var v389 int32
	_ = v389
	var v390 int64
	_ = v390
	var v391 int32
	_ = v391
	var v392 int64
	_ = v392
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v398 int64
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v409 int64
	_ = v409
	var v414 int64
	_ = v414
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int64
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v461 int64
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int64
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int64
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int64
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int64
	_ = v499
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int64
	_ = v519
	var v520 int64
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int64
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v539 int64
	_ = v539
	var v540 int32
	_ = v540
	var v541 int64
	_ = v541
	var v542 int32
	_ = v542
	var v543 int64
	_ = v543
	var v544 int32
	_ = v544
	var v545 int64
	_ = v545
	var v546 int32
	_ = v546
	var v547 int64
	_ = v547
	var v551 int64
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v562 int64
	_ = v562
	var v571 int64
	_ = v571
	var v583 int64
	_ = v583
	v9 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = v10 & l3
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v12&l3 != 0 {
		if v11 != 0 {
			return int32(0)
		} else {
			v17 = l1 + int32(8)
			v18 = int32(3)
			if v18 < l2 {
				v261 = int64(0)
				if base.B2i32(v17 != (l1+int32(11))&int32(-4))|base.B2i32(l2 < int32(4)) != 0 {
					v340 = v17
					v341 = l2
					v346 = v261
				} else {
					v271 = l2 - int32(4)
					v275 = int32(base.Ui32(v271)>>(uint(int32(2))%32)) + int32(1)
					v277 = v275 & int32(3)
					if base.Ui32(int32(12)) <= base.Ui32(v271) {
						v282 = v17
						v283 = l2
						v286 = int32(0)
						v288 = v261
						for {
							v289 = int32(16)
							v290 = v283 - v289
							v292 = v282 + v289
							v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
							v296 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
							v299 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
							v302 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
							v308 = base.I64_extend_i32_u(base.I32_popcnt(v293)) + (base.I64_extend_i32_u(base.I32_popcnt(v296)) + (base.I64_extend_i32_u(base.I32_popcnt(v299)) + (v288 + base.I64_extend_i32_u(base.I32_popcnt(v302)))))
							v310 = v286 + int32(4)
							if v310 != v275&int32(2147483644) {
								v282 = v292
								v283 = v290
								v286 = v310
								v288 = v308
								continue
							} else {
								break
							}
							break
						}
						if v277 == int32(0) {
							v340 = v292
							v341 = v290
							v346 = v308
						} else {
							v314 = v292
							v315 = v290
							v320 = v308
							v322 = v314
							v323 = v315
							v324 = int32(0)
							v328 = v320
							for {
								v329 = int32(4)
								v330 = v323 - v329
								v332 = v322 + v329
								v333 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
								v336 = v328 + base.I64_extend_i32_u(base.I32_popcnt(v333))
								v338 = v324 + int32(1)
								if v338 != v277 {
									v322 = v332
									v323 = v330
									v324 = v338
									v328 = v336
									continue
								} else {
									break
								}
								break
							}
							v340 = v332
							v341 = v330
							v346 = v336
						}
					} else {
						v314 = v17
						v315 = l2
						v320 = v261
						v322 = v314
						v323 = v315
						v324 = int32(0)
						v328 = v320
						for {
							v329 = int32(4)
							v330 = v323 - v329
							v332 = v322 + v329
							v333 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
							v336 = v328 + base.I64_extend_i32_u(base.I32_popcnt(v333))
							v338 = v324 + int32(1)
							if v338 != v277 {
								v322 = v332
								v323 = v330
								v324 = v338
								v328 = v336
								continue
							} else {
								break
							}
							break
						}
						v340 = v332
						v341 = v330
						v346 = v336
					}
				}
				if v341 == int32(0) {
					v409 = v346
				} else {
					v350 = v341 & int32(3)
					if v350 == int32(0) {
						v371 = v340
						v373 = v341
						v377 = v346
					} else {
						v354 = v340
						v356 = v341
						v358 = int32(0)
						v360 = v346
						for {
							v361 = int32(1)
							v362 = v354 + v361
							v364 = v356 - v361
							v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
							v366 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v365)+uint32(_c_Fn13925[0]))))
							v367 = v360 + v366
							v369 = v358 + v361
							if v369 != v350 {
								v354 = v362
								v356 = v364
								v358 = v369
								v360 = v367
								continue
							} else {
								break
							}
							break
						}
						v371 = v362
						v373 = v364
						v377 = v367
					}
					if base.Ui32(v341) < base.Ui32(int32(4)) {
						v409 = v377
					} else {
						v380 = v371
						v382 = v373
						v386 = v377
						for {
							v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+3)))
							v388 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v387)+uint32(_c_Fn13925[0]))))
							v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+2)))
							v390 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v389)+uint32(_c_Fn13925[0]))))
							v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+1)))
							v392 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v391)+uint32(_c_Fn13925[0]))))
							v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
							v394 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v393)+uint32(_c_Fn13925[0]))))
							v398 = v388 + (v390 + (v392 + (v386 + v394)))
							v399 = int32(4)
							v402 = v382 - v399
							if v402 != 0 {
								v380 = v380 + v399
								v382 = v402
								v386 = v398
								continue
							} else {
								break
							}
							break
						}
						v409 = v398
					}
				}
				v583 = v409
			} else {
				if l2 == int32(0) {
					v583 = v9
				} else {
					v25 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v31 = int32(0)
						v32 = v17
						v39 = v9
						for {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+3)))
							v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v40)+uint32(_c_Fn13925[0]))))
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+2)))
							v47 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v44)+uint32(_c_Fn13925[0]))))
							v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
							v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_Fn13925[0]))))
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
							v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_c_Fn13925[0]))))
							v59 = v43 + (v47 + (v51 + (v39 + v55)))
							v60 = int32(4)
							v61 = v32 + v60
							v63 = v31 + v60
							if v63 != l2&int32(-4) {
								v31 = v63
								v32 = v61
								v39 = v59
								continue
							} else {
								break
							}
							break
						}
						if v25 == int32(0) {
							v583 = v59
						} else {
							v68 = v61
							v75 = v59
							v77 = int32(0)
							v78 = v68
							v85 = v75
							for {
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
								v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v86)+uint32(_c_Fn13925[0]))))
								v90 = v85 + v89
								v91 = int32(1)
								v94 = v77 + v91
								if v94 != v25 {
									v77 = v94
									v78 = v78 + v91
									v85 = v90
									continue
								} else {
									break
								}
								break
							}
							v583 = v90
						}
					} else {
						v68 = v17
						v75 = v9
						v77 = int32(0)
						v78 = v68
						v85 = v75
						for {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
							v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v86)+uint32(_c_Fn13925[0]))))
							v90 = v85 + v89
							v91 = int32(1)
							v94 = v77 + v91
							if v94 != v25 {
								v77 = v94
								v78 = v78 + v91
								v85 = v90
								continue
							} else {
								break
							}
							break
						}
						v583 = v90
					}
				}
			}
			return l2<<(uint(v18)%32) - base.I32_wrap_i64(v583)
		}
	} else {
		if v11 != 0 {
			v97 = l0 + int32(8)
			v98 = int32(3)
			if v98 < l2 {
				v414 = int64(0)
				if base.B2i32(v97 != (l0+int32(11))&int32(-4))|base.B2i32(l2 < int32(4)) != 0 {
					v493 = v97
					v494 = l2
					v499 = v414
				} else {
					v424 = l2 - int32(4)
					v428 = int32(base.Ui32(v424)>>(uint(int32(2))%32)) + int32(1)
					v430 = v428 & int32(3)
					if base.Ui32(int32(12)) <= base.Ui32(v424) {
						v435 = v97
						v436 = l2
						v439 = int32(0)
						v441 = v414
						for {
							v442 = int32(16)
							v443 = v436 - v442
							v445 = v435 + v442
							v446 = *(*int32)(unsafe.Add(mBase, uint32(v435)+12))
							v449 = *(*int32)(unsafe.Add(mBase, uint32(v435)+8))
							v452 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
							v455 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
							v461 = base.I64_extend_i32_u(base.I32_popcnt(v446)) + (base.I64_extend_i32_u(base.I32_popcnt(v449)) + (base.I64_extend_i32_u(base.I32_popcnt(v452)) + (v441 + base.I64_extend_i32_u(base.I32_popcnt(v455)))))
							v463 = v439 + int32(4)
							if v463 != v428&int32(2147483644) {
								v435 = v445
								v436 = v443
								v439 = v463
								v441 = v461
								continue
							} else {
								break
							}
							break
						}
						if v430 == int32(0) {
							v493 = v445
							v494 = v443
							v499 = v461
						} else {
							v467 = v445
							v468 = v443
							v473 = v461
							v475 = v467
							v476 = v468
							v477 = int32(0)
							v481 = v473
							for {
								v482 = int32(4)
								v483 = v476 - v482
								v485 = v475 + v482
								v486 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
								v489 = v481 + base.I64_extend_i32_u(base.I32_popcnt(v486))
								v491 = v477 + int32(1)
								if v491 != v430 {
									v475 = v485
									v476 = v483
									v477 = v491
									v481 = v489
									continue
								} else {
									break
								}
								break
							}
							v493 = v485
							v494 = v483
							v499 = v489
						}
					} else {
						v467 = v97
						v468 = l2
						v473 = v414
						v475 = v467
						v476 = v468
						v477 = int32(0)
						v481 = v473
						for {
							v482 = int32(4)
							v483 = v476 - v482
							v485 = v475 + v482
							v486 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
							v489 = v481 + base.I64_extend_i32_u(base.I32_popcnt(v486))
							v491 = v477 + int32(1)
							if v491 != v430 {
								v475 = v485
								v476 = v483
								v477 = v491
								v481 = v489
								continue
							} else {
								break
							}
							break
						}
						v493 = v485
						v494 = v483
						v499 = v489
					}
				}
				if v494 == int32(0) {
					v562 = v499
				} else {
					v503 = v494 & int32(3)
					if v503 == int32(0) {
						v524 = v493
						v526 = v494
						v530 = v499
					} else {
						v507 = v493
						v509 = v494
						v511 = int32(0)
						v513 = v499
						for {
							v514 = int32(1)
							v515 = v507 + v514
							v517 = v509 - v514
							v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
							v519 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v518)+uint32(_c_Fn13925[0]))))
							v520 = v513 + v519
							v522 = v511 + v514
							if v522 != v503 {
								v507 = v515
								v509 = v517
								v511 = v522
								v513 = v520
								continue
							} else {
								break
							}
							break
						}
						v524 = v515
						v526 = v517
						v530 = v520
					}
					if base.Ui32(v494) < base.Ui32(int32(4)) {
						v562 = v530
					} else {
						v533 = v524
						v535 = v526
						v539 = v530
						for {
							v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+3)))
							v541 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v540)+uint32(_c_Fn13925[0]))))
							v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+2)))
							v543 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v542)+uint32(_c_Fn13925[0]))))
							v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+1)))
							v545 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v544)+uint32(_c_Fn13925[0]))))
							v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
							v547 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v546)+uint32(_c_Fn13925[0]))))
							v551 = v541 + (v543 + (v545 + (v539 + v547)))
							v552 = int32(4)
							v555 = v535 - v552
							if v555 != 0 {
								v533 = v533 + v552
								v535 = v555
								v539 = v551
								continue
							} else {
								break
							}
							break
						}
						v562 = v551
					}
				}
				v571 = v562
			} else {
				if l2 == int32(0) {
					v571 = v9
				} else {
					v105 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v111 = int32(0)
						v112 = v97
						v119 = v9
						for {
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+3)))
							v123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_Fn13925[0]))))
							v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+2)))
							v127 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v124)+uint32(_c_Fn13925[0]))))
							v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
							v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v128)+uint32(_c_Fn13925[0]))))
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
							v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_c_Fn13925[0]))))
							v139 = v123 + (v127 + (v131 + (v119 + v135)))
							v140 = int32(4)
							v141 = v112 + v140
							v143 = v111 + v140
							if v143 != l2&int32(-4) {
								v111 = v143
								v112 = v141
								v119 = v139
								continue
							} else {
								break
							}
							break
						}
						if v105 == int32(0) {
							v571 = v139
						} else {
							v148 = v141
							v155 = v139
							v157 = int32(0)
							v158 = v148
							v165 = v155
							for {
								v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
								v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v166)+uint32(_c_Fn13925[0]))))
								v170 = v165 + v169
								v171 = int32(1)
								v174 = v157 + v171
								if v174 != v105 {
									v157 = v174
									v158 = v158 + v171
									v165 = v170
									continue
								} else {
									break
								}
								break
							}
							v571 = v170
						}
					} else {
						v148 = v97
						v155 = v9
						v157 = int32(0)
						v158 = v148
						v165 = v155
						for {
							v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v166)+uint32(_c_Fn13925[0]))))
							v170 = v165 + v169
							v171 = int32(1)
							v174 = v157 + v171
							if v174 != v105 {
								v157 = v174
								v158 = v158 + v171
								v165 = v170
								continue
							} else {
								break
							}
							break
						}
						v571 = v170
					}
				}
			}
			return l2<<(uint(v98)%32) - base.I32_wrap_i64(v571)
		} else {
			if l2 <= int32(0) {
				return int32(0)
			} else {
				v180 = int32(8)
				v181 = l1 + v180
				v183 = l0 + v180
				v184 = int32(0)
				if l2 != int32(1) {
					v193 = v184
					v194 = v184
					v195 = int32(0)
					for {
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v183))))
						v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v181))))
						v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203^v205)+uint32(_c_Fn13925[0]))))
						v212 = v194 | int32(1)
						v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v212))))
						v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212+v183))))
						v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214^v216)+uint32(_c_Fn13925[0]))))
						v221 = v193 + v209 + v220
						v222 = int32(2)
						v223 = v194 + v222
						v225 = v195 + v222
						if v225 != l2&int32(2147483646) {
							v193 = v221
							v194 = v223
							v195 = v225
							continue
						} else {
							break
						}
						break
					}
					if l2&int32(1) == int32(0) {
						v247 = v221
					} else {
						v229 = v221
						v230 = v223
						v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v181))))
						v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v183))))
						v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239^v241)+uint32(_c_Fn13925[0]))))
						v247 = v229 + v245
					}
				} else {
					v229 = v184
					v230 = v184
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v181))))
					v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v183))))
					v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239^v241)+uint32(_c_Fn13925[0]))))
					v247 = v229 + v245
				}
				return v247
			}
		}
	}
}
func Fn13936(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_JsonbExtractScalar(m, v13+int32(4), v10+int32(12))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			if v21 != 0 {
				switch v23 {
				case 0:
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v24 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v41 = int32(0)
							m.G0 = v10 + int32(32)
							return v41
						}
					} else {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						v41 = int32(0)
						m.G0 = v10 + int32(32)
						return v41
					}
				default:
					F_cannotCastJsonbValue(m, v23, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
					v35 = F_DirectFunctionCall1Coll(m, l2, int32(0), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v13 == v37 {
							v41 = v35
							m.G0 = v10 + int32(32)
							return v41
						} else {
							F_pfree(m, v13)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v35
								m.G0 = v10 + int32(32)
								return v41
							}
						}
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v23, l1)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func Fn13938(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, l3, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v20 = F_local2local(m, v9, v8, v13, l3, l2, l1, base.B2i32(v10 != int32(0)))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return v20
		}
	}
}
func Fn13945(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_palloc0(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l3
		return v7
	}
}
func Fn13952(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v10 != 0 {
		v13 = int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = v12
	}
	v16 = F_numeric_poly_stddev_internal(m, v13, l2, l1, v8+int32(15))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v20 == int32(1) {
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
			v26 = int32(0)
		} else {
			v26 = v16
		}
		m.G0 = v8 + int32(16)
		return v26
	}
}
func Fn13954(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
		v16 = F_pg_snprintf(m, v11, l2, l1, v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v11
		}
	}
}
func Fn13965(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
		if v20 == int32(_a_Fn13965_0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, l4, int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_Fn13965_1), l3, l2)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v14
			v38 = v11 + int32(16)
			v41 = F_pg_snprintf(m, v38, int32(32), int32(_a_Fn13965_2), v11)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v44 = int32(0)
				v50 = F_DirectFunctionCall3Coll(m, int32(408), v44, v38, v44, int32(-1))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = F_DirectFunctionCall2Coll(m, l1, v44, v50, v16)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = F_DirectFunctionCall1Coll(m, int32(1465), v44, v52)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(48)
							return v54
						}
					}
				}
			}
		}
	}
}
func Fn13972(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v16 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v20 = F_pg_detoast_datum_packed(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = v20
					v23 = F_encrypt_internal(m, l2, l1, v9, v14, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v25 != v9 {
							F_pfree(m, v9)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v29 != v14 {
									F_pfree(m, v14)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int32(0)
									} else {
										v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v33 < int32(3) {
											return v23
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v22 == v36 {
												return v23
											} else {
												F_pfree(m, v22)
												mBase = m.M
												v39 = m.ExcPending
												if v39 != 0 {
													return int32(0)
												} else {
													return v23
												}
											}
										}
									}
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v22 == v36 {
											return v23
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v22 == v36 {
											return v23
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v22 == v36 {
										return v23
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					}
				}
			} else {
				v22 = int32(0)
				v23 = F_encrypt_internal(m, l2, l1, v9, v14, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v25 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v22 == v36 {
											return v23
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v22 == v36 {
										return v23
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v29 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v22 == v36 {
										return v23
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						} else {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v33 < int32(3) {
								return v23
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v22 == v36 {
									return v23
								} else {
									F_pfree(m, v22)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										return v23
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
func Fn13978(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v7 <= l3 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
		switch v9 - int32(3) {
		case 0, 2:
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v15 = v14
		default:
			v15 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
	} else {
	}
	return int32(0)
}
func Fn13981(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v6 = F_palloc0(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v14
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
		return v14
	}
}
func Fn13990(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = F_patternsel_common(m, v3, v4, v5, v6, v7, v8, l1, v5)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func Fn13996(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v4 = l3
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v17 != 0 {
			v18 = F_array_contains_nulls(m, v13)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_Fn13996_0), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13996_1), l2, l1)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v23 = F_ArrayGetNItemsSafe(m, v20, v13+int32(16))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v4)
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						if v26 != 0 {
							v34 = v26
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v34 = (v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v34+v13, v23, v10+int32(15))
						mBase = m.M
						m.G0 = v10 + int32(16)
						return v13
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v23 = F_ArrayGetNItemsSafe(m, v20, v13+int32(16))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v4)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				if v26 != 0 {
					v34 = v26
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v34 = (v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				F_isort(m, v34+v13, v23, v10+int32(15))
				mBase = m.M
				m.G0 = v10 + int32(16)
				return v13
			}
		}
	}
}
func Fn14005(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v10 = m.G0
	v11 = int32(-64)
	v12 = v10 + v11
	m.G0 = v12
	v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+20)))
	v16 = v12 - v11
	v17 = v16
	v21 = v14
	goto L1
L1:
	;
	v27 = v17 - int32(1)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v21)&l3)+uint32(_c_Fn14005[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v30)
	if base.Ui64(v21) < base.Ui64(l2) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v36 = v16 - v27
	v38 = v36 + int32(4)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v12) < base.Ui32(v27) {
		v17 = v27
		v21 = int64(base.Ui64(v21) >> (uint(l1) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(int32(2)) % 32)
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	base.MemoryCopy(m, v39+int32(4), v27, v36)
	goto L10
L9:
	;
	goto L10
L10:
	;
	m.G0 = v12 - int32(-64)
	return v39
}
func Fn14009(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = l1
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_palloc0(m, int32(36))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(440)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l7
		v27 = F_list_make1_impl(m, int32(472), v13+int32(8))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l6
			*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l2
			*(*uint16)(unsafe.Add(mBase, uint32(v16)+8)) = uint16(v2)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v27
			m.G0 = v13 + int32(16)
			return v16
		}
	}
}
func Fn14012(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v24) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return int32(0)
L5:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errmsg(m, int32(_a_Fn14012_0), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(_a_Fn14012_1), l2, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	return v144
L10:
	;
	v144 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_Fn14012[0]))
	if v35 == v24 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v144 = int32(1)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_Fn14012[1]))
	if v39 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v144 = v136
	goto L9
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_Fn14012[2]))
	if v43 == int32(0) {
		v136 = int32(0)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_Fn14012[3]))
	v107 = int32(0)
	v109 = v39 - int32(1)
	goto L39
L20:
	;
	v48 = v43
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v53 == int32(4) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v136 = int32(0)
	goto L16
L23:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	if v100 != 0 {
		v48 = v100
		goto L21
	} else {
		goto L38
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v56 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v59 = int32(1)
	if v24 == v56 {
		v136 = v59
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v63 = v61 - int32(1)
	if v63 < int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v68 = int32(0)
	v70 = v63
	goto L28
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v76 = int32(2)
	v77 = base.I32_div_s(v70-v68, v76)
	v78 = v77 + v68
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74+v78<<(uint(v76)%32))))
	if v82 == v24 {
		v136 = v59
		goto L16
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v86 = F_TransactionIdPrecedes(m, v82, v24)
	mBase = m.M
	if v86 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v87 = v78 + int32(1)
	goto L33
L32:
	;
	v87 = v68
	goto L33
L33:
	;
	if v86 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v90 = v70
	goto L36
L35:
	;
	v90 = v78 - int32(1)
	goto L36
L36:
	;
	if v87 <= v90 {
		v68 = v87
		v70 = v90
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L29
L38:
	;
	goto L22
L39:
	;
	v114 = int32(2)
	v115 = base.I32_div_s(v109-v107, v114)
	v116 = v115 + v107
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v105+v116<<(uint(v114)%32))))
	v121 = base.B2i32(v120 == v24)
	if v120 == v24 {
		v136 = v121
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v136 = v121
	goto L16
L41:
	;
	v124 = base.B2i32(base.Ui32(v120) < base.Ui32(v24))
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v125 = v116 + int32(1)
	goto L44
L43:
	;
	v125 = v107
	goto L44
L44:
	;
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v128 = v109
	goto L47
L46:
	;
	v128 = v116 - int32(1)
	goto L47
L47:
	;
	if v125 <= v128 {
		v107 = v125
		v109 = v128
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
}
func Fn14014(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_SearchSysCache1(m, l5, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				F_errmsg_internal(m, l4, v11)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+84))
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(16)
				return v31
			}
		}
	}
}
func Fn14018(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+29)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v20
	v24 = F_DirectFunctionCall1Coll(m, int32(3376), int32(0), v10)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(48)
		return v24
	}
}
func Fn14025(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 float32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 float64
	_ = v98
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v18 = v11 + int32(1)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v24 = v22 & int32(1)
			if v24 != 0 {
				v25 = v18
			} else {
				v25 = v11 + int32(4)
			}
			if v22 == int32(1) {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v31 == int32(18) {
					v34 = int32(16)
				} else {
					v34 = int32(0)
				}
				if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v41 = int32(4)
				} else {
					v41 = v34
				}
				v52 = v41
			} else {
				v42 = int32(1)
				if v24 != 0 {
					v52 = int32(base.Ui32(v22)>>(uint(v42)%32)) - v42
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v53 = int32(1)
			v54 = v20 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v20 + int32(4)
			}
			if v57 == int32(1) {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v66 == int32(18) {
					v69 = int32(16)
				} else {
					v69 = int32(0)
				}
				if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v76 = int32(4)
				} else {
					v76 = v69
				}
				v87 = v76
			} else {
				v77 = int32(1)
				if v59 != 0 {
					v87 = int32(base.Ui32(v57)>>(uint(v77)%32)) - v77
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					v87 = int32(base.Ui32(v81)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v88 = F_calc_word_similarity(m, v25, v52, v60, v87, l2)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v90 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v94 != v20 {
							F_pfree(m, v20)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								return base.F64_le(v98, base.F64_promote_f32(v88))
							}
						} else {
							v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.F64_le(v98, base.F64_promote_f32(v88))
						}
					}
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v94 != v20 {
						F_pfree(m, v20)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.F64_le(v98, base.F64_promote_f32(v88))
						}
					} else {
						v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						return base.F64_le(v98, base.F64_promote_f32(v88))
					}
				}
			}
		}
	}
}
