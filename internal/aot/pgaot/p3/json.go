package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonTableDestroyOpaque(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v3 = F_GetJsonTableExecContext(m, l0, int32(_a_F_JsonTableDestroyOpaque_0))
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = v5
		*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v5
		return
	}
}
func F_JsonTableGetValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	v8 = F_GetJsonTableExecContext(m, l0, int32(_a_F_JsonTableGetValue_0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v14 = l1 << (uint(int32(2)) % 32)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v14)))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+40)))
		if v17 == int32(1) {
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v20)
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v25+v14)))
			if v27 != 0 {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+44)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
				v31 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v28)+44)) = uint8(v31)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v30
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
				v36 = m.T0[v35].(func(*base.Module, int32, int32, int32) int32)(m, v27, v28, l4)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+44)) = uint8(v29)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v33
					return v36
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
				v42 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v42)
				return v41
			}
		}
	}
}
func F_appendJSONKeyValueFmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_appendJSONKeyValueFmt[0]))
	v17 = F_palloc(m, int32(128))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_appendJSONKeyValueFmt[0])) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
	v23 = F_pvsnprintf(m, v17, int32(128), l3, l4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v23) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = v17
	v33 = v23
	goto L7
L5:
	;
	v51 = v17
	goto L6
L6:
	;
	if v51 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	F_pfree(m, v32)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v51 = v38
	goto L6
L9:
	;
	v38 = F_palloc(m, v33)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_appendJSONKeyValueFmt[0])) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
	v43 = F_pvsnprintf(m, v38, v33, l3, l4)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if base.Ui32(v33) <= base.Ui32(v43) {
		v32 = v38
		v33 = v43
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	F_pfree(m, v51)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_escape_json(m, l0, l1)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_appendStringInfoChar(m, l0, int32(58))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if l2 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_escape_json(m, l0, v51)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_appendStringInfoString(m, l0, v51)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L13
L22:
	;
	goto L13
L23:
	;
	m.G0 = v12 + int32(16)
	return
}
func F_json_agg_transfn(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_json_agg_transfn_worker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_json_categorize_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v28 int32
	_ = v28
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_getBaseType(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
		if v10 <= int32(1081) {
			switch v10 - int32(16) {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1243)
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
				m.G0 = v8 + int32(16)
				return
			case 1, 2, 3, 6:
				v74 = F_get_element_type(m, v10)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					if v74 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
						m.G0 = v8 + int32(16)
						return
					} else {
						switch v10 - int32(2277) {
						case 0, 10:
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
							m.G0 = v8 + int32(16)
							return
						case 1, 2, 3, 4, 5, 6, 7, 8, 9:
							v84 = F_type_is_rowtype(m, v10)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								if v84 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
									m.G0 = v8 + int32(16)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
									if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
										v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											if v98 != int32(1) {
												F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
													return
												}
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												if v102 == int32(0) {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
													m.G0 = v8 + int32(16)
													return
												}
											}
										}
									} else {
										F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						default:
							if v10 != int32(_a_F_json_categorize_type_1) {
								v84 = F_type_is_rowtype(m, v10)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									if v84 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
										m.G0 = v8 + int32(16)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
										if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
											v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												if v98 != int32(1) {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												} else {
													v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
													if v102 == int32(0) {
														F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															m.G0 = v8 + int32(16)
															return
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
														*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
														m.G0 = v8 + int32(16)
														return
													}
												}
											}
										} else {
											F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												m.G0 = v8 + int32(16)
												return
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			case 4, 5, 7:
				F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2)
					m.G0 = v8 + int32(16)
					return
				}
			default:
				if base.Ui32(v10-int32(700)) < base.Ui32(int32(2)) {
					F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2)
						m.G0 = v8 + int32(16)
						return
					}
				} else {
					if v10 != int32(114) {
						v74 = F_get_element_type(m, v10)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							if v74 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
								m.G0 = v8 + int32(16)
								return
							} else {
								switch v10 - int32(2277) {
								case 0, 10:
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
									m.G0 = v8 + int32(16)
									return
								case 1, 2, 3, 4, 5, 6, 7, 8, 9:
									v84 = F_type_is_rowtype(m, v10)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										if v84 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
											m.G0 = v8 + int32(16)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
											if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
												v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													if v98 != int32(1) {
														F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															m.G0 = v8 + int32(16)
															return
														}
													} else {
														v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
														if v102 == int32(0) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
															*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
															m.G0 = v8 + int32(16)
															return
														}
													}
												}
											} else {
												F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
													return
												}
											}
										}
									}
								default:
									if v10 != int32(_a_F_json_categorize_type_1) {
										v84 = F_type_is_rowtype(m, v10)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											if v84 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
												m.G0 = v8 + int32(16)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
												if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
													v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														if v98 != int32(1) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
															if v102 == int32(0) {
																F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(16)
																	return
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
																*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
																m.G0 = v8 + int32(16)
																return
															}
														}
													}
												} else {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					} else {
						F_getTypeOutputInfo(m, int32(114), l3, v8+int32(15))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(6)
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			if v10 <= int32(1183) {
				if v10 == int32(1082) {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1085)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(3)
					m.G0 = v8 + int32(16)
					return
				} else {
					if v10 != int32(1114) {
						v74 = F_get_element_type(m, v10)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							if v74 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
								m.G0 = v8 + int32(16)
								return
							} else {
								switch v10 - int32(2277) {
								case 0, 10:
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
									m.G0 = v8 + int32(16)
									return
								case 1, 2, 3, 4, 5, 6, 7, 8, 9:
									v84 = F_type_is_rowtype(m, v10)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										if v84 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
											m.G0 = v8 + int32(16)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
											if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
												v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													if v98 != int32(1) {
														F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															m.G0 = v8 + int32(16)
															return
														}
													} else {
														v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
														if v102 == int32(0) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
															*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
															m.G0 = v8 + int32(16)
															return
														}
													}
												}
											} else {
												F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
													return
												}
											}
										}
									}
								default:
									if v10 != int32(_a_F_json_categorize_type_1) {
										v84 = F_type_is_rowtype(m, v10)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											if v84 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
												m.G0 = v8 + int32(16)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
												if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
													v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														if v98 != int32(1) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
															if v102 == int32(0) {
																F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(16)
																	return
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
																*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
																m.G0 = v8 + int32(16)
																return
															}
														}
													}
												} else {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1313)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(4)
						m.G0 = v8 + int32(16)
						return
					}
				}
			} else {
				if v10 == int32(1184) {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1151)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(5)
					m.G0 = v8 + int32(16)
					return
				} else {
					if v10 == int32(1700) {
						F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2)
							m.G0 = v8 + int32(16)
							return
						}
					} else {
						if v10 != int32(3802) {
							v74 = F_get_element_type(m, v10)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								if v74 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
									m.G0 = v8 + int32(16)
									return
								} else {
									switch v10 - int32(2277) {
									case 0, 10:
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
										m.G0 = v8 + int32(16)
										return
									case 1, 2, 3, 4, 5, 6, 7, 8, 9:
										v84 = F_type_is_rowtype(m, v10)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											if v84 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
												m.G0 = v8 + int32(16)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
												if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
													v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														if v98 != int32(1) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
															if v102 == int32(0) {
																F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(16)
																	return
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
																*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
																m.G0 = v8 + int32(16)
																return
															}
														}
													}
												} else {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												}
											}
										}
									default:
										if v10 != int32(_a_F_json_categorize_type_1) {
											v84 = F_type_is_rowtype(m, v10)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												if v84 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
													m.G0 = v8 + int32(16)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
													if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
														v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return
														} else {
															if v98 != int32(1) {
																F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(16)
																	return
																}
															} else {
																v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
																if v102 == int32(0) {
																	F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																	mBase = m.M
																	v112 = m.ExcPending
																	if v112 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(16)
																		return
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
																	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
																	m.G0 = v8 + int32(16)
																	return
																}
															}
														}
													} else {
														F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															m.G0 = v8 + int32(16)
															return
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						} else {
							F_getTypeOutputInfo(m, int32(3802), l3, v8+int32(15))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								if l1 != 0 {
									v54 = int32(7)
								} else {
									v54 = int32(6)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v54
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_json_manifest_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v10 - int32(3) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = l1
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = F_strtox_2(m, l1, v6+int32(-4), int32(10), int64(-9223372036854775807-1))
		mBase = m.M
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
		if v21 != 0 {
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_json_manifest_scalar_0)
			m.T0[v86].(func(*base.Module, int32, int32, int32))(m, v85, int32(_a_F_json_manifest_scalar_1), v6+int32(-32))
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v22 = base.I32_wrap_i64(v19)
			if base.Ui32(v22-int32(3)) <= base.Ui32(int32(-3)) {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a_F_json_manifest_scalar_2)
				m.T0[v95].(func(*base.Module, int32, int32, int32))(m, v94, int32(_a_F_json_manifest_scalar_1), v6+int32(-48))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				m.T0[v27].(func(*base.Module, int32, int32))(m, v14, v22)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
					m.G0 = v8 - int32(-64)
					return int32(0)
				}
			}
		}
	case 1:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = l1
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v40 = F_strtox_2(m, l1, v6+int32(-4), int32(10), int64(-1))
		mBase = m.M
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
		if v42 != 0 {
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(_a_F_json_manifest_scalar_3)
			m.T0[v104].(func(*base.Module, int32, int32, int32))(m, v103, int32(_a_F_json_manifest_scalar_1), v6+int32(-16))
			mBase = m.M
			v111 = m.ExcPending
			if v111 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
			m.T0[v43].(func(*base.Module, int32, int64))(m, v35, v40)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
				m.G0 = v8 - int32(-64)
				return int32(0)
			}
		}
	default:
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_json_manifest_scalar_4)
		m.T0[v68].(func(*base.Module, int32, int32, int32))(m, v67, int32(_a_F_json_manifest_scalar_1), v8)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 5:
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		switch v48 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		case 1:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		case 2:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		case 3:
			F_pfree(m, l1)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
				m.G0 = v8 - int32(-64)
				return int32(0)
			}
		case 4:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		case 5:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		default:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		}
	case 9:
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if base.Ui32(v58) <= base.Ui32(int32(2)) {
			*(*int32)(unsafe.Add(mBase, uint32(l0+v58<<(uint(int32(2))%32))+40)) = l1
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(11)
		m.G0 = v8 - int32(-64)
		return int32(0)
	case 10:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
		m.G0 = v8 - int32(-64)
		return int32(0)
	}
}
func F_json_object_agg_unique_transfn(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_json_object_agg_transfn_worker(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_json_parse_manifest_incremental_chunk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v397 int32
	_ = v397
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v620 int32
	_ = v620
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v740 int64
	_ = v740
	var v741 int64
	_ = v741
	var v743 int64
	_ = v743
	var v744 int64
	_ = v744
	var v747 int64
	_ = v747
	var v750 int64
	_ = v750
	var v752 int64
	_ = v752
	var v753 int64
	_ = v753
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v800 int32
	_ = v800
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	v4 = l3
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if l0 == int32(_a_F_json_parse_manifest_incremental_chunk_0) {
		v429 = int32(16)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v429 == v4^int32(1) {
		goto L122
	} else {
		goto L123
	}
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v29 == int32(_a_F_json_parse_manifest_incremental_chunk_1) {
		v429 = int32(16)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v33 != int32(1) {
		v429 = int32(2)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = l0 + int32(68)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)) = uint8(v4)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v45)
	v47 = F_json_lex(m, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v429 = v397
	goto L1
L6:
	;
	return
L7:
	;
	if v47 != 0 {
		v397 = v47
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v50 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v397 = int32(0)
	goto L5
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v54 = int32(_a_F_json_parse_manifest_incremental_chunk_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v53))) = uint16(v54)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v58 = v56 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v58
	if v58 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v62 = v50
	goto L12
L12:
	;
	v67 = v62
	v71 = v49
	goto L14
L13:
	;
	v62 = v58
	goto L12
L14:
	;
	v82 = v67 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v85 = v84 + v82
	v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85))))
	if v86 == v71 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L9
L16:
	;
	if v369 != 0 {
		v67 = v369
		v71 = v371
		goto L14
	} else {
		goto L121
	}
L17:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v369 = v368
	v371 = v365
	goto L16
L18:
	;
	if base.Ui32(int32(11)) < base.Ui32(v71) {
		v365 = v71
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v86&int32(32) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v90 = F_json_lex(m, l0)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	if v90 != 0 {
		v397 = v90
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v365 = v92
	goto L17
L24:
	;
	v117 = v86 & int32(255)
	if v86&int32(64) != 0 {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v101 = v86*int32(104) + v71<<(uint(int32(3))%32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_json_parse_manifest_incremental_chunk[0])))
	if v104 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_json_parse_manifest_incremental_chunk[1])))
	if v109 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	base.MemoryCopy(m, v85, v104, v109)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v112 = v111 + v109
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v112
	v369 = v112
	v371 = v71
	goto L16
L30:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v339&int32(4) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L31:
	;
	switch v117 + int32(-64) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	case 3:
		goto L41
	case 4:
		goto L40
	case 5:
		goto L39
	case 6:
		goto L38
	case 7:
		goto L37
	case 8:
		goto L36
	case 9:
		goto L35
	case 10:
		goto L34
	default:
		v365 = v71
		goto L17
	}
L32:
	;
	goto L33
L33:
	;
	switch v117 - int32(1) {
	case 0:
		goto L111
	default:
		goto L104
	case 3, 35:
		goto L110
	case 5, 33:
		v327 = int32(3)
		goto L103
	case 6:
		goto L109
	case 7:
		goto L108
	case 11:
		goto L107
	case 32:
		goto L106
	case 34:
		goto L105
	}
L34:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	if v282 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L96
	}
L35:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	v247 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v247
	if v246 == v247 {
		v365 = v71
		goto L17
	} else {
		goto L83
	}
L36:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	if v233 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L80
	}
L37:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v224 = base.B2i32(v71 == int32(11))
	*(*uint8)(unsafe.Add(mBase, uint32(v220+v221))) = uint8(v224)
	if v218 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L77
	}
L38:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	if v200 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L74
	}
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v184 = base.B2i32(v71 == int32(11))
	*(*uint8)(unsafe.Add(mBase, uint32(v180+v181))) = uint8(v184)
	if v178 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L71
	}
L40:
	;
	v164 = int32(0)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	if v165|v166 == v164 {
		v338 = v164
		goto L30
	} else {
		goto L67
	}
L41:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	F_dec_lex_level(m, l0)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L63
	}
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(_a_F_json_parse_manifest_incremental_chunk_3) < v144 {
		v429 = int32(3)
		goto L1
	} else {
		goto L56
	}
L43:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	F_dec_lex_level(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L52
	}
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(_a_F_json_parse_manifest_incremental_chunk_3) < v123 {
		v429 = int32(3)
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v126 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v128 = m.T0[v126].(func(*base.Module, int32) int32)(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_inc_lex_level(m, l0)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	if v128 != 0 {
		v397 = v128
		goto L5
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v365 = v71
	goto L17
L52:
	;
	if v133 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L53
	}
L53:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v139 = m.T0[v133].(func(*base.Module, int32) int32)(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	if v139 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L55
	}
L55:
	;
	v397 = v139
	goto L5
L56:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v147 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v149 = m.T0[v147].(func(*base.Module, int32) int32)(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_inc_lex_level(m, l0)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	if v149 != 0 {
		v397 = v149
		goto L5
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v365 = v71
	goto L17
L63:
	;
	if v154 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L64
	}
L64:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v160 = m.T0[v154].(func(*base.Module, int32) int32)(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	if v160 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L66
	}
L66:
	;
	v397 = v160
	goto L5
L67:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v170 != int32(1) {
		v338 = v164
		goto L30
	} else {
		goto L68
	}
L68:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = F_pstrdup(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	if v175 != 0 {
		v338 = v175
		goto L30
	} else {
		goto L70
	}
L70:
	;
	v429 = int32(16)
	goto L1
L71:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190+v191<<(uint(int32(2))%32))))
	v196 = m.T0[v178].(func(*base.Module, int32, int32, int32) int32)(m, v188, v195, v184)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	if v196 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L73
	}
L73:
	;
	v397 = v196
	goto L5
L74:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(int32(2))%32))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+v206))))
	v214 = m.T0[v200].(func(*base.Module, int32, int32, int32) int32)(m, v203, v210, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	if v214 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L76
	}
L76:
	;
	v397 = v214
	goto L5
L77:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v229 = m.T0[v218].(func(*base.Module, int32, int32) int32)(m, v228, v224)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	if v229 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L79
	}
L79:
	;
	v397 = v229
	goto L5
L80:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+16))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+v239))))
	v242 = m.T0[v233].(func(*base.Module, int32, int32) int32)(m, v236, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	if v242 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L82
	}
L82:
	;
	v397 = v242
	goto L5
L83:
	;
	if v71 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v71
	v365 = v71
	goto L17
L85:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v253 != int32(1) {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v264 = v262 - v263
	v267 = F_palloc(m, v264+int32(1))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L6
	} else {
		goto L91
	}
L88:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v258 = F_pstrdup(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v258
	if v258 != 0 {
		goto L84
	} else {
		goto L90
	}
L90:
	;
	v429 = int32(16)
	goto L1
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v267
	if v267 == int32(0) {
		v429 = int32(16)
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v264 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	base.MemoryCopy(m, v267, v273, v264)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v277 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v275+v264))) = uint8(v277)
	goto L84
L96:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v288 = m.T0[v282].(func(*base.Module, int32, int32, int32) int32)(m, v285, v286, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v290&int32(4) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v301 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v301
	if v288 == v301 {
		v365 = v71
		goto L17
	} else {
		goto L102
	}
L99:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	if v295 == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	F_pfree(m, v295)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	v397 = v288
	goto L5
L103:
	;
	v328 = int32(11)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v329 == int32(0) {
		v397 = v328
		goto L5
	} else {
		goto L115
	}
L104:
	;
	v327 = int32(0)
	goto L103
L105:
	;
	v327 = int32(4)
	goto L103
L106:
	;
	v327 = int32(2)
	goto L103
L107:
	;
	v327 = int32(8)
	goto L103
L108:
	;
	v327 = int32(5)
	goto L103
L109:
	;
	v316 = int32(1)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85-v316))))
	if v318 == v316 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v327 = int32(6)
	goto L103
L111:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85-int32(1)))))
	v327 = base.B2i32(v310 == int32(8))
	goto L103
L112:
	;
	v321 = int32(6)
	goto L114
L113:
	;
	v321 = int32(3)
	goto L114
L114:
	;
	v327 = v321
	goto L103
L115:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v332 == int32(12) {
		v397 = v328
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v327<<(uint(int32(2))%32))+uint32(_c_F_json_parse_manifest_incremental_chunk[2])))
	v429 = v337
	goto L1
L117:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v357+v358<<(uint(int32(2))%32)))) = v338
	v365 = v71
	goto L17
L118:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v345+v346<<(uint(int32(2))%32))))
	if v350 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	F_pfree(m, v350)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	goto L15
L122:
	;
	if v4 != 0 {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	goto L124
L124:
	;
	v822 = F_json_errdetail(m, v429, l0)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L6
	} else {
		goto L218
	}
L125:
	;
	m.G0 = v21 + int32(32)
	return
L126:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v452 = int32(0)
	v453 = m.G0
	v455 = v453 - int32(112)
	m.G0 = v455
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if l2 == v452 {
		goto L136
	} else {
		goto L137
	}
L127:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v433 == int32(14) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v443 = F_pg_cryptohash_update(m, v442, l1, l2)
	mBase = m.M
	if int32(0) <= v443 {
		goto L125
	} else {
		goto L132
	}
L130:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(_a_F_json_parse_manifest_incremental_chunk_4)
	m.T0[v436].(func(*base.Module, int32, int32, int32))(m, v24, int32(_a_F_json_parse_manifest_incremental_chunk_5), v21)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	m.T0[v448].(func(*base.Module, int32, int32, int32))(m, v24, int32(_a_F_json_parse_manifest_incremental_chunk_6), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	goto L125
L134:
	;
	goto L125
L135:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+32)) = int32(_a_F_json_parse_manifest_incremental_chunk_7)
	m.T0[v793].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_5), v455+int32(32))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L6
	} else {
		goto L217
	}
L136:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = int32(_a_F_json_parse_manifest_incremental_chunk_8)
	m.T0[v787].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_5), v455)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L216
	}
L137:
	;
	if l2 != int32(1) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if base.Ui32(v548) <= base.Ui32(int32(1)) {
		goto L136
	} else {
		goto L164
	}
L139:
	;
	v471 = v452
	v475 = v5
	v476 = v5
	v477 = v5
	v482 = v5
	goto L142
L140:
	;
	v516 = v5
	v517 = v5
	v518 = v5
	v523 = v5
	goto L141
L141:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v517))))
	v528 = base.B2i32(v526 == int32(10))
	if v526 == int32(10) {
		goto L158
	} else {
		goto L159
	}
L142:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v476))))
	v487 = base.B2i32(v485 == int32(10))
	if v485 == int32(10) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if l2&int32(1) == int32(0) {
		v541 = v497
		v543 = v496
		v548 = v499
		goto L138
	} else {
		goto L157
	}
L144:
	;
	v488 = v476
	goto L146
L145:
	;
	v488 = v475
	goto L146
L146:
	;
	if v485 == int32(10) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v489 = v475
	goto L149
L148:
	;
	v489 = v477
	goto L149
L149:
	;
	v491 = v476 | int32(1)
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v491))))
	v495 = base.B2i32(v493 == int32(10))
	if v493 == int32(10) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v496 = v488
	goto L152
L151:
	;
	v496 = v489
	goto L152
L152:
	;
	if v493 == int32(10) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v497 = v491
	goto L155
L154:
	;
	v497 = v488
	goto L155
L155:
	;
	v499 = v487 + v482 + v495
	v500 = int32(2)
	v501 = v476 + v500
	v503 = v471 + v500
	if v503 != l2&int32(-2) {
		v471 = v503
		v475 = v497
		v476 = v501
		v477 = v496
		v482 = v499
		goto L142
	} else {
		goto L156
	}
L156:
	;
	goto L143
L157:
	;
	v516 = v497
	v517 = v501
	v518 = v496
	v523 = v499
	goto L141
L158:
	;
	v529 = v516
	goto L160
L159:
	;
	v529 = v518
	goto L160
L160:
	;
	if v526 == int32(10) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v530 = v517
	goto L163
L162:
	;
	v530 = v516
	goto L163
L163:
	;
	v541 = v530
	v543 = v529
	v548 = v528 + v523
	goto L138
L164:
	;
	if v541 != l2-int32(1) {
		goto L135
	} else {
		goto L165
	}
L165:
	;
	if v451 != 0 {
		v575 = v451
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v576 = F_pg_cryptohash_update(m, v575, l1, v543+int32(1))
	mBase = m.M
	if v576 < int32(0) {
		goto L175
	} else {
		goto L176
	}
L167:
	;
	v558 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	if v558 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v564].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_9), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L6
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v567 = F_pg_cryptohash_init(m, v558)
	mBase = m.M
	if int32(0) <= v567 {
		v575 = v558
		goto L166
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v572].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_10), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	v575 = v558
	goto L166
L175:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v581].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_6), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L6
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v587 = F_pg_cryptohash_final(m, v575, v455+int32(80), int32(32))
	mBase = m.M
	if v587 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v592].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_11), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L6
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	if v595 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L181
L183:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v601].(func(*base.Module, int32, int32, int32))(m, v598, int32(_a_F_json_parse_manifest_incremental_chunk_12), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L6
	} else {
		goto L186
	}
L184:
	;
	v605 = v595
	goto L185
L185:
	;
	v606 = F_strlen(m, v605)
	mBase = m.M
	if v606 != int32(64) {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v605 = v604
	goto L185
L187:
	;
	v740 = *(*int64)(unsafe.Add(mBase, uint32(v455)+80))
	v741 = *(*int64)(unsafe.Add(mBase, uint32(v455)+48))
	v743 = *(*int64)(unsafe.Add(mBase, uint32(v455)+88))
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v455)+56))
	v747 = *(*int64)(unsafe.Add(mBase, uint32(v455)+96))
	v750 = *(*int64)(unsafe.Add(mBase, uint32(v455-int32(-64))))
	v752 = *(*int64)(unsafe.Add(mBase, uint32(v455)+104))
	v753 = *(*int64)(unsafe.Add(mBase, uint32(v455)+72))
	if v740^v741|(v743^v744)|(v747^v750|(v752^v753)) != int64(0) {
		goto L211
	} else {
		goto L212
	}
L188:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+16)) = v605
	m.T0[v715].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_13), v455+int32(16))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L210
	}
L189:
	;
	v620 = int32(0)
	goto L190
L190:
	;
	v630 = v605 + v620<<(uint(int32(1))%32)
	v631 = int32(*(*int8)(unsafe.Add(mBase, uint32(v630))))
	v633 = v631 - int32(48)
	if base.Ui32(v633&int32(255)) <= base.Ui32(int32(9)) {
		v656 = v633
		goto L192
	} else {
		goto L193
	}
L191:
	;
	goto L187
L192:
	;
	v657 = int32(*(*int8)(unsafe.Add(mBase, uint32(v630)+1)))
	v659 = v657 - int32(48)
	if base.Ui32(v659&int32(255)) <= base.Ui32(int32(9)) {
		v682 = v659
		goto L200
	} else {
		goto L201
	}
L193:
	;
	if base.Ui32((v631-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v656 = v631 - int32(87)
	goto L192
L195:
	;
	goto L196
L196:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v631-int32(65))&int32(255)) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v655 = int32(-1)
	goto L199
L198:
	;
	v655 = v631 - int32(55)
	goto L199
L199:
	;
	v656 = v655
	goto L192
L200:
	;
	if v682|v656 < int32(0) {
		goto L188
	} else {
		goto L208
	}
L201:
	;
	if base.Ui32((v657-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v682 = v657 - int32(87)
	goto L200
L203:
	;
	goto L204
L204:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v657-int32(65))&int32(255)) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v681 = int32(-1)
	goto L207
L206:
	;
	v681 = v657 - int32(55)
	goto L207
L207:
	;
	v682 = v681
	goto L200
L208:
	;
	v691 = v682 + v656<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v455+int32(48)+v620))) = uint8(v691)
	v694 = v620 + int32(1)
	if v694 != int32(32) {
		v620 = v694
		goto L190
	} else {
		goto L209
	}
L209:
	;
	goto L191
L210:
	;
	goto L187
L211:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v761].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_14), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L6
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	F_pg_cryptohash_free(m, v575)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L6
	} else {
		goto L215
	}
L214:
	;
	goto L213
L215:
	;
	m.G0 = v455 + int32(112)
	goto L134
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v822
	m.T0[v824].(func(*base.Module, int32, int32, int32))(m, v24, int32(_a_F_json_parse_manifest_incremental_chunk_5), v21+int32(16))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L6
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_json_populate_record(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = int32(1)
	v6 = F_populate_record_worker(m, l0, int32(_a_F_json_populate_record_0), v3, v3, int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_json_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v9 = v6 + int32(8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v16 = F_pq_getmsgtext(m, v10, v11-v12, v6+int32(76))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_json_recv[0]))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
		v25 = F_makeJsonLexContextCstringLen(m, v9, v16, v20, v23, int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v29 = F_pg_parse_json_or_errsave(m, v9, int32(_a_F_json_recv_0), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
				v32 = F_cstring_to_text_with_len(m, v16, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(80)
					return v32
				}
			}
		}
	}
}
func F_makeJsonLexContextCstringLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v5 = l4
	if l0 == int32(0) {
		v9 = F_palloc0(m, int32(68))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				return int32(_a_F_makeJsonLexContextCstringLen_0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v17 | int32(1)
				v24 = v9
				*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v24))) = l1
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+56)) = uint8(v5)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l2
				if v5 != 0 {
					v35 = F_makeStringInfo(m)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v35
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v38 | int32(2)
						return v24
					}
				} else {
					return v24
				}
			}
		}
	} else {
		base.MemoryFill(m, l0, int32(0), int32(68))
		v24 = l0
		*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v24))) = l1
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+56)) = uint8(v5)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l2
		if v5 != 0 {
			v35 = F_makeStringInfo(m)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v35
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v38 | int32(2)
				return v24
			}
		} else {
			return v24
		}
	}
}
func F_printJsonPathItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	F_check_stack_depth(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_printJsonPathItem[0]))
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v18 {
	case 0:
		goto L8
	case 1:
		goto L47
	case 2:
		goto L46
	case 3:
		goto L45
	case 4, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 41:
		goto L44
	case 6:
		goto L43
	case 7:
		goto L42
	case 19, 20:
		goto L41
	case 21:
		goto L40
	case 22:
		goto L39
	case 23:
		goto L38
	case 24:
		goto L37
	case 25:
		goto L36
	case 26:
		goto L35
	case 27:
		goto L34
	case 28:
		goto L33
	case 29:
		goto L32
	case 30:
		goto L31
	case 31:
		goto L30
	case 32:
		goto L29
	case 33:
		goto L28
	case 34:
		goto L27
	case 35:
		goto L26
	case 36:
		goto L25
	case 37:
		goto L24
	case 38:
		goto L23
	default:
		goto L9
	case 40:
		goto L22
	case 42:
		goto L21
	case 43:
		goto L20
	case 44:
		goto L19
	case 45:
		goto L18
	case 46:
		goto L17
	case 47:
		goto L16
	case 48:
		goto L15
	case 49:
		goto L14
	case 50:
		goto L13
	case 51:
		goto L12
	case 52:
		goto L11
	case 53:
		goto L10
	}
L6:
	;
	goto L5
L7:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v691 {
		goto L296
	} else {
		goto L297
	}
L8:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L295
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L292
	}
L10:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_1))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L285
	}
L11:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_2))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L278
	}
L12:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_3))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L271
	}
L13:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_4))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L264
	}
L14:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_5))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L263
	}
L15:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_6))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L262
	}
L16:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_7))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L261
	}
L17:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_8))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L248
	}
L18:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_9))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L247
	}
L19:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_10))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L246
	}
L20:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_11))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L245
	}
L21:
	;
	if l3 != 0 {
		goto L202
	} else {
		goto L203
	}
L22:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_12))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L201
	}
L23:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_13))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L200
	}
L24:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_14))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L193
	}
L25:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_15))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L192
	}
L26:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_16))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L191
	}
L27:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_17))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L190
	}
L28:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_18))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L189
	}
L29:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_19))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L188
	}
L30:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_20))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L187
	}
L31:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_21))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L183
	}
L32:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_22))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L179
	}
L33:
	;
	F_appendStringInfoChar(m, l0, int32(36))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L177
	}
L34:
	;
	F_appendStringInfoChar(m, l0, int32(36))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L176
	}
L35:
	;
	F_appendStringInfoChar(m, l0, int32(64))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L175
	}
L36:
	;
	if l2 != 0 {
		goto L170
	} else {
		goto L171
	}
L37:
	;
	if l2 != 0 {
		goto L147
	} else {
		goto L148
	}
L38:
	;
	F_appendStringInfoChar(m, l0, int32(91))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L120
	}
L39:
	;
	if l2 != 0 {
		goto L115
	} else {
		goto L116
	}
L40:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_23))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L114
	}
L41:
	;
	if l3 != 0 {
		goto L96
	} else {
		goto L97
	}
L42:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L92
	}
L43:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_24))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L88
	}
L44:
	;
	if l3 != 0 {
		goto L62
	} else {
		goto L63
	}
L45:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v43 != 0 {
		goto L57
	} else {
		goto L58
	}
L46:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v23 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_escape_json_with_len(m, l0, v19, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L7
L49:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	F_appendStringInfoString(m, l0, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v36 <= int32(0) {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L7
L57:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_25))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_26))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L7
L61:
	;
	goto L7
L62:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v55, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v62 = v60 - int32(4)
	if base.Ui32(v62) <= base.Ui32(int32(37)) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v68 = v67
	goto L69
L68:
	;
	v68 = int32(6)
	goto L69
L69:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v74 = v72 - int32(4)
	if base.Ui32(v74) <= base.Ui32(int32(37)) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v81 = v79
	goto L72
L71:
	;
	v81 = int32(6)
	goto L72
L72:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v68) <= base.Ui32(v81)))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_appendStringInfoChar(m, l0, int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v89 = F_jspOperationName(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_appendStringInfoString(m, l0, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_appendStringInfoChar(m, l0, int32(32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_jspInitByBuffer(m, v10+int32(100), v98, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v105 = v103 - int32(4)
	if base.Ui32(v105) <= base.Ui32(int32(37)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v111 = v110
	goto L81
L80:
	;
	v111 = int32(6)
	goto L81
L81:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v117 = v115 - int32(4)
	if base.Ui32(v117) <= base.Ui32(int32(37)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v124 = v122
	goto L84
L83:
	;
	v124 = int32(6)
	goto L84
L84:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v111) <= base.Ui32(v124)))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L7
L88:
	;
	v137 = v10 + int32(100)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v137, v138, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v142 = int32(0)
	F_printJsonPathItem(m, l0, v137, v142, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	goto L7
L92:
	;
	v153 = v10 + int32(100)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v153, v154, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v158 = int32(0)
	F_printJsonPathItem(m, l0, v153, v158, v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_27))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L7
L96:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	v171 = v18
	goto L98
L98:
	;
	if v171 == int32(19) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v171 = v170
	goto L98
L100:
	;
	v174 = int32(43)
	goto L102
L101:
	;
	v174 = int32(45)
	goto L102
L102:
	;
	F_appendStringInfoChar(m, l0, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v179, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v186 = v184 - int32(4)
	if base.Ui32(v186) <= base.Ui32(int32(37)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v186<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v192 = v191
	goto L107
L106:
	;
	v192 = int32(6)
	goto L107
L107:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v198 = v196 - int32(4)
	if base.Ui32(v198) <= base.Ui32(int32(37)) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v205 = v203
	goto L110
L109:
	;
	v205 = int32(6)
	goto L110
L110:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v192) <= base.Ui32(v205)))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	goto L7
L114:
	;
	goto L7
L115:
	;
	F_appendStringInfoChar(m, l0, int32(46))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	F_appendStringInfoChar(m, l0, int32(42))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	goto L7
L120:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v226 <= int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L146
	}
L122:
	;
	v230 = v10 + int32(100)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	F_jspInitByBuffer(m, v230, v231, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v237 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v239 = v10 + int32(72)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v239, v240, v237)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	v253 = v10 + int32(100)
	goto L126
L126:
	;
	v254 = int32(0)
	F_printJsonPathItem(m, l0, v253, v254, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L130
	}
L127:
	;
	v243 = int32(0)
	F_printJsonPathItem(m, l0, v230, v243, v243)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_28))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v253 = v239
	goto L126
L130:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v258 < int32(2) {
		goto L121
	} else {
		goto L131
	}
L131:
	;
	v265 = int32(1)
	goto L132
L132:
	;
	v270 = v10 + int32(100)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v273 = v265 << (uint(int32(3)) % 32)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273+v274)))
	F_jspInitByBuffer(m, v270, v271, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L134
	}
L133:
	;
	goto L121
L134:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v279+v273)+4))
	if v281 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v304 = int32(0)
	F_printJsonPathItem(m, l0, v303, v304, v304)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L144
	}
L136:
	;
	v283 = v10 + int32(72)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v283, v284, v281)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L143
	}
L139:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v290 = int32(0)
	F_printJsonPathItem(m, l0, v270, v290, v290)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_28))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v303 = v283
	goto L135
L143:
	;
	v303 = v10 + int32(100)
	goto L135
L144:
	;
	v309 = v265 + int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v309 < v310 {
		v265 = v309
		goto L132
	} else {
		goto L145
	}
L145:
	;
	goto L133
L146:
	;
	goto L7
L147:
	;
	F_appendStringInfoChar(m, l0, int32(46))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v326 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	goto L149
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v326
	F_appendStringInfo(m, l0, int32(_a_F_printJsonPathItem_29), v10+int32(16))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L169
	}
L152:
	;
	if v326 == int32(-1) {
		goto L163
	} else {
		goto L164
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v326
	F_appendStringInfo(m, l0, int32(_a_F_printJsonPathItem_30), v10+int32(32))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L162
	}
L154:
	;
	switch v325 + int32(1) {
	case 0:
		goto L157
	case 1:
		goto L153
	default:
		goto L151
	}
L155:
	;
	goto L156
L156:
	;
	if v326 != v325 {
		goto L152
	} else {
		goto L159
	}
L157:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_31))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	goto L7
L159:
	;
	if v326 != int32(-1) {
		goto L153
	} else {
		goto L160
	}
L160:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_32))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	goto L7
L162:
	;
	goto L7
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v325
	F_appendStringInfo(m, l0, int32(_a_F_printJsonPathItem_33), v10+int32(48))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	if v325 != int32(-1) {
		goto L151
	} else {
		goto L167
	}
L166:
	;
	goto L7
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v326
	F_appendStringInfo(m, l0, int32(_a_F_printJsonPathItem_34), v10-int32(-64))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	goto L7
L169:
	;
	goto L7
L170:
	;
	F_appendStringInfoChar(m, l0, int32(46))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_escape_json_with_len(m, l0, v372, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L174
	}
L173:
	;
	goto L172
L174:
	;
	goto L7
L175:
	;
	goto L7
L176:
	;
	goto L7
L177:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_escape_json_with_len(m, l0, v385, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L7
L179:
	;
	v393 = v10 + int32(100)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v393, v394, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v398 = int32(0)
	F_printJsonPathItem(m, l0, v393, v398, v398)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	goto L7
L183:
	;
	v409 = v10 + int32(100)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v409, v410, v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v414 = int32(0)
	F_printJsonPathItem(m, l0, v409, v414, v414)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	goto L7
L187:
	;
	goto L7
L188:
	;
	goto L7
L189:
	;
	goto L7
L190:
	;
	goto L7
L191:
	;
	goto L7
L192:
	;
	goto L7
L193:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v442 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v444 = v10 + int32(100)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v444, v445, v442)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L199
	}
L197:
	;
	v448 = int32(0)
	F_printJsonPathItem(m, l0, v444, v448, v448)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	goto L7
L200:
	;
	goto L7
L201:
	;
	goto L7
L202:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v466 = v10 + int32(100)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v466, v467, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L206
	}
L205:
	;
	goto L204
L206:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v474 = v472 - int32(4)
	if base.Ui32(v474) <= base.Ui32(int32(37)) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v484 = v482 - int32(4)
	if base.Ui32(v484) <= base.Ui32(int32(37)) {
		goto L212
	} else {
		goto L213
	}
L208:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v474<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v481 = v479
	goto L210
L209:
	;
	v481 = int32(6)
	goto L210
L210:
	;
	goto L207
L211:
	;
	F_printJsonPathItem(m, l0, v466, int32(0), base.B2i32(base.Ui32(v481) <= base.Ui32(v491)))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L215
	}
L212:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v484<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v491 = v489
	goto L214
L213:
	;
	v491 = int32(6)
	goto L214
L214:
	;
	goto L211
L215:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_35))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_escape_json_with_len(m, l0, v498, v499)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v502 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_36))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L243
	}
L221:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v506&int32(1) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	F_appendStringInfoChar(m, l0, int32(105))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	v513 = v506
	goto L224
L224:
	;
	if v513&int32(2) != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v513 = v512
	goto L224
L226:
	;
	F_appendStringInfoChar(m, l0, int32(115))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	v520 = v513
	goto L228
L228:
	;
	if v520&int32(4) != 0 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v520 = v519
	goto L228
L230:
	;
	F_appendStringInfoChar(m, l0, int32(109))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	v527 = v520
	goto L232
L232:
	;
	if v527&int32(8) != 0 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v527 = v526
	goto L232
L234:
	;
	F_appendStringInfoChar(m, l0, int32(120))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	v534 = v527
	goto L236
L236:
	;
	if v534&int32(16) != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v534 = v533
	goto L236
L238:
	;
	F_appendStringInfoChar(m, l0, int32(113))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L242
	}
L241:
	;
	goto L240
L242:
	;
	goto L220
L243:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	goto L7
L245:
	;
	goto L7
L246:
	;
	goto L7
L247:
	;
	goto L7
L248:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v561 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v563 = v10 + int32(100)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v563, v564, v561)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v572 != 0 {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	v567 = int32(0)
	F_printJsonPathItem(m, l0, v563, v567, v567)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	goto L251
L254:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L260
	}
L257:
	;
	v577 = v10 + int32(100)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_jspInitByBuffer(m, v577, v578, v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v582 = int32(0)
	F_printJsonPathItem(m, l0, v577, v582, v582)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	goto L256
L260:
	;
	goto L7
L261:
	;
	goto L7
L262:
	;
	goto L7
L263:
	;
	goto L7
L264:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v602 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v604 = v10 + int32(100)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v604, v605, v602)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	v608 = int32(0)
	F_printJsonPathItem(m, l0, v604, v608, v608)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	goto L7
L271:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v619 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v621 = v10 + int32(100)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v621, v622, v619)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	v625 = int32(0)
	F_printJsonPathItem(m, l0, v621, v625, v625)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	goto L274
L277:
	;
	goto L7
L278:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v636 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v638 = v10 + int32(100)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v638, v639, v636)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L284
	}
L282:
	;
	v642 = int32(0)
	F_printJsonPathItem(m, l0, v638, v642, v642)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	goto L7
L285:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v653 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v655 = v10 + int32(100)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v655, v656, v653)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	v659 = int32(0)
	F_printJsonPathItem(m, l0, v655, v659, v659)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	goto L288
L291:
	;
	goto L7
L292:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v671
	F_errmsg_internal(m, int32(_a_F_printJsonPathItem_37), v10)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_printJsonPathItem_38), int32(835), int32(_a_F_printJsonPathItem_39))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	goto L7
L296:
	;
	v695 = v10 + int32(100)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v695, v696, v691)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	m.G0 = v10 + int32(128)
	return
L299:
	;
	v699 = int32(1)
	F_printJsonPathItem(m, l0, v695, v699, v699)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	goto L298
}
