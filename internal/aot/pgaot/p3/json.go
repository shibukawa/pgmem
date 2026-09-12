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
	v3 = F_GetJsonTableExecContext(m, l0, int32(361422))
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
	v8 = F_GetJsonTableExecContext(m, l0, int32(364063))
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[43]))
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
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v15
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
	*(*int32)(unsafe.Add(mBase, _consts[43])) = v15
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
									if base.Ui32(int32(16384)) <= base.Ui32(v10) {
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
							if v10 != int32(5078) {
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
										if base.Ui32(int32(16384)) <= base.Ui32(v10) {
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
											if base.Ui32(int32(16384)) <= base.Ui32(v10) {
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
									if v10 != int32(5078) {
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
												if base.Ui32(int32(16384)) <= base.Ui32(v10) {
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
											if base.Ui32(int32(16384)) <= base.Ui32(v10) {
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
									if v10 != int32(5078) {
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
												if base.Ui32(int32(16384)) <= base.Ui32(v10) {
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
												if base.Ui32(int32(16384)) <= base.Ui32(v10) {
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
										if v10 != int32(5078) {
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
													if base.Ui32(int32(16384)) <= base.Ui32(v10) {
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
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(235710)
			m.T0[v86].(func(*base.Module, int32, int32, int32))(m, v85, int32(209701), v6+int32(-32))
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
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(283781)
				m.T0[v95].(func(*base.Module, int32, int32, int32))(m, v94, int32(209701), v6+int32(-48))
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
			*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(235610)
			m.T0[v104].(func(*base.Module, int32, int32, int32))(m, v103, int32(209701), v6+int32(-16))
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
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(240968)
		m.T0[v68].(func(*base.Module, int32, int32, int32))(m, v67, int32(209701), v8)
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
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
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v401 int32
	_ = v401
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
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
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v607 int32
	_ = v607
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v837 int32
	_ = v837
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	v4 = l3
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v26 = l0 + int32(68)
	if l0 == int32(4555704) {
		v433 = int32(16)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v433 == v4^int32(1) {
		goto L124
	} else {
		goto L125
	}
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v31 == int32(4555772) {
		v433 = int32(16)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v35 != int32(1) {
		v433 = int32(2)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)) = uint8(v4)
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
	v433 = v401
	goto L1
L6:
	;
	return
L7:
	;
	if v47 != 0 {
		v401 = v47
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
	v401 = int32(0)
	goto L5
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v54 = int32(8204)
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
	if v373 != 0 {
		v67 = v373
		v71 = v375
		goto L14
	} else {
		goto L123
	}
L17:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v373 = v372
	v375 = v369
	goto L16
L18:
	;
	if base.Ui32(int32(11)) < base.Ui32(v71) {
		v369 = v71
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
		v401 = v90
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v369 = v92
	goto L17
L24:
	;
	v118 = v86 & int32(255)
	if v86&int32(64) != 0 {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v101 = v86*int32(104) + v71<<(uint(int32(3))%32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_consts[1238])))
	if v104 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_consts[1239])))
	if v109 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v113 = v109 + v112
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v113
	v373 = v113
	v375 = v71
	goto L16
L28:
	;
	v110 = F__emscripten_memcpy_bulkmem(m, v85, v104, v109)
	mBase = m.M
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v343&int32(4) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L32:
	;
	switch v118 + int32(-64) {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L43
	case 3:
		goto L42
	case 4:
		goto L41
	case 5:
		goto L40
	case 6:
		goto L39
	case 7:
		goto L38
	case 8:
		goto L37
	case 9:
		goto L36
	case 10:
		goto L35
	default:
		v369 = v71
		goto L17
	}
L33:
	;
	goto L34
L34:
	;
	switch v118 - int32(1) {
	case 0:
		goto L113
	default:
		v329 = int32(0)
		goto L105
	case 3, 35:
		goto L111
	case 5, 33:
		goto L112
	case 6:
		goto L110
	case 7:
		goto L109
	case 11:
		goto L108
	case 32:
		goto L107
	case 34:
		goto L106
	}
L35:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v284 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L98
	}
L36:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v248 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v248
	if v247 == v248 {
		v369 = v71
		goto L17
	} else {
		goto L84
	}
L37:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v234 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L81
	}
L38:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v225 = base.B2i32(v71 == int32(11))
	*(*uint8)(unsafe.Add(mBase, uint32(v221+v222))) = uint8(v225)
	if v219 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L78
	}
L39:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v201 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L75
	}
L40:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v185 = base.B2i32(v71 == int32(11))
	*(*uint8)(unsafe.Add(mBase, uint32(v181+v182))) = uint8(v185)
	if v179 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L72
	}
L41:
	;
	v165 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v166|v167 == v165 {
		v342 = v165
		goto L31
	} else {
		goto L68
	}
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	F_dec_lex_level(m, l0)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L64
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(6399) < v145 {
		v433 = int32(3)
		goto L1
	} else {
		goto L57
	}
L44:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	F_dec_lex_level(m, l0)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L6
	} else {
		goto L53
	}
L45:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(6399) < v124 {
		v433 = int32(3)
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v127 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v129 = m.T0[v127].(func(*base.Module, int32) int32)(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_inc_lex_level(m, l0)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	if v129 != 0 {
		v401 = v129
		goto L5
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v369 = v71
	goto L17
L53:
	;
	if v134 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v140 = m.T0[v134].(func(*base.Module, int32) int32)(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	if v140 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L56
	}
L56:
	;
	v401 = v140
	goto L5
L57:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v148 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v150 = m.T0[v148].(func(*base.Module, int32) int32)(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	F_inc_lex_level(m, l0)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L63
	}
L61:
	;
	if v150 != 0 {
		v401 = v150
		goto L5
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v369 = v71
	goto L17
L64:
	;
	if v155 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L65
	}
L65:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v161 = m.T0[v155].(func(*base.Module, int32) int32)(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	if v161 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L67
	}
L67:
	;
	v401 = v161
	goto L5
L68:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v171 != int32(1) {
		v342 = v165
		goto L31
	} else {
		goto L69
	}
L69:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = F_pstrdup(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	if v176 != 0 {
		v342 = v176
		goto L31
	} else {
		goto L71
	}
L71:
	;
	v433 = int32(16)
	goto L1
L72:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v191+v192<<(uint(int32(2))%32))))
	v197 = m.T0[v179].(func(*base.Module, int32, int32, int32) int32)(m, v189, v196, v185)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	if v197 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L74
	}
L74:
	;
	v401 = v197
	goto L5
L75:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v206+v207<<(uint(int32(2))%32))))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v205)+16))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212+v207))))
	v215 = m.T0[v201].(func(*base.Module, int32, int32, int32) int32)(m, v204, v211, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	if v215 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L77
	}
L77:
	;
	v401 = v215
	goto L5
L78:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v230 = m.T0[v219].(func(*base.Module, int32, int32) int32)(m, v229, v225)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	if v230 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L80
	}
L80:
	;
	v401 = v230
	goto L5
L81:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v240))))
	v243 = m.T0[v234].(func(*base.Module, int32, int32) int32)(m, v237, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	if v243 == int32(0) {
		v369 = v71
		goto L17
	} else {
		goto L83
	}
L83:
	;
	v401 = v243
	goto L5
L84:
	;
	if v71 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v71
	v369 = v71
	goto L17
L86:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v254 != int32(1) {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v265 = v263 - v264
	v268 = F_palloc(m, v265+int32(1))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L92
	}
L89:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v259 = F_pstrdup(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v259
	if v259 != 0 {
		goto L85
	} else {
		goto L91
	}
L91:
	;
	v433 = int32(16)
	goto L1
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v268
	if v268 == int32(0) {
		v433 = int32(16)
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v265 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v279 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v277+v265))) = uint8(v279)
	goto L85
L95:
	;
	v275 = F__emscripten_memcpy_bulkmem(m, v268, v274, v265)
	mBase = m.M
	goto L97
L96:
	;
	goto L97
L97:
	;
	goto L94
L98:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v290 = m.T0[v284].(func(*base.Module, int32, int32, int32) int32)(m, v287, v288, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v292&int32(4) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v303 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v303
	if v290 == v303 {
		v369 = v71
		goto L17
	} else {
		goto L104
	}
L101:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	if v297 == int32(0) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	F_pfree(m, v297)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	v401 = v290
	goto L5
L105:
	;
	v330 = int32(11)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v331 == int32(0) {
		v401 = v330
		goto L5
	} else {
		goto L117
	}
L106:
	;
	v329 = int32(4)
	goto L105
L107:
	;
	v329 = int32(2)
	goto L105
L108:
	;
	v329 = int32(8)
	goto L105
L109:
	;
	v329 = int32(5)
	goto L105
L110:
	;
	v319 = int32(1)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85-v319))))
	if v321 == v319 {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v329 = int32(6)
	goto L105
L112:
	;
	v329 = int32(3)
	goto L105
L113:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85-int32(1)))))
	v329 = base.B2i32(v312 == int32(8))
	goto L105
L114:
	;
	v324 = int32(6)
	goto L116
L115:
	;
	v324 = int32(3)
	goto L116
L116:
	;
	v329 = v324
	goto L105
L117:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v334 == int32(12) {
		v401 = v330
		goto L5
	} else {
		goto L118
	}
L118:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v329<<(uint(int32(2))%32))+uint32(_consts[1240])))
	v433 = v341
	goto L1
L119:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v361+v362<<(uint(int32(2))%32)))) = v342
	v369 = v71
	goto L17
L120:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v349+v350<<(uint(int32(2))%32))))
	if v354 == int32(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	F_pfree(m, v354)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	goto L119
L123:
	;
	goto L15
L124:
	;
	if v4 != 0 {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	goto L126
L126:
	;
	v859 = F_json_errdetail(m, v433, l0)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L6
	} else {
		goto L240
	}
L127:
	;
	m.G0 = v21 + int32(32)
	return
L128:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v456 = int32(0)
	v457 = m.G0
	v459 = v457 - int32(112)
	m.G0 = v459
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if l2 == v456 {
		goto L138
	} else {
		goto L139
	}
L129:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v437 == int32(14) {
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v447 = F_pg_cryptohash_update(m, v446, l1, l2)
	mBase = m.M
	if int32(0) <= v447 {
		goto L127
	} else {
		goto L134
	}
L132:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(20846)
	m.T0[v440].(func(*base.Module, int32, int32, int32))(m, v24, int32(209701), v21)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	m.T0[v452].(func(*base.Module, int32, int32, int32))(m, v24, int32(83608), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	goto L127
L136:
	;
	goto L127
L137:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v459)+32)) = int32(468336)
	m.T0[v830].(func(*base.Module, int32, int32, int32))(m, v461, int32(209701), v459+int32(32))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L6
	} else {
		goto L239
	}
L138:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = int32(172670)
	m.T0[v824].(func(*base.Module, int32, int32, int32))(m, v461, int32(209701), v459)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L6
	} else {
		goto L238
	}
L139:
	;
	v464 = int32(1)
	if l2 != v464 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v479 = v5
	v480 = v5
	v482 = v456
	v483 = v5
	v486 = v5
	goto L143
L141:
	;
	v518 = v5
	v519 = v5
	v522 = v5
	v525 = v5
	goto L142
L142:
	;
	if l2&v464 != 0 {
		goto L158
	} else {
		goto L159
	}
L143:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v479))))
	v491 = base.B2i32(v489 == int32(10))
	if v489 == int32(10) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v518 = v505
	v519 = v501
	v522 = v500
	v525 = v503
	goto L142
L145:
	;
	v492 = v479
	goto L147
L146:
	;
	v492 = v480
	goto L147
L147:
	;
	if v489 == int32(10) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v493 = v480
	goto L150
L149:
	;
	v493 = v483
	goto L150
L150:
	;
	v495 = v479 | int32(1)
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v495))))
	v499 = base.B2i32(v497 == int32(10))
	if v497 == int32(10) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v500 = v492
	goto L153
L152:
	;
	v500 = v493
	goto L153
L153:
	;
	if v497 == int32(10) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v501 = v495
	goto L156
L155:
	;
	v501 = v492
	goto L156
L156:
	;
	v503 = v491 + v486 + v499
	v504 = int32(2)
	v505 = v479 + v504
	v507 = v482 + v504
	if v507 != l2&int32(-2) {
		v479 = v505
		v480 = v501
		v482 = v507
		v483 = v500
		v486 = v503
		goto L143
	} else {
		goto L157
	}
L157:
	;
	goto L144
L158:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v518))))
	v530 = base.B2i32(v528 == int32(10))
	if v528 == int32(10) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v535 = v519
	v536 = v522
	v537 = v525
	goto L160
L160:
	;
	if base.Ui32(v537) <= base.Ui32(int32(1)) {
		goto L138
	} else {
		goto L167
	}
L161:
	;
	v531 = v519
	goto L163
L162:
	;
	v531 = v522
	goto L163
L163:
	;
	if v528 == int32(10) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v532 = v518
	goto L166
L165:
	;
	v532 = v519
	goto L166
L166:
	;
	v535 = v532
	v536 = v531
	v537 = v530 + v525
	goto L160
L167:
	;
	if v535 != l2-int32(1) {
		goto L137
	} else {
		goto L168
	}
L168:
	;
	if v455 != 0 {
		v563 = v455
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v564 = F_pg_cryptohash_update(m, v563, l1, v536+int32(1))
	mBase = m.M
	if v564 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L170:
	;
	v546 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L6
	} else {
		goto L171
	}
L171:
	;
	if v546 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	m.T0[v552].(func(*base.Module, int32, int32, int32))(m, v461, int32(14020), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L6
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v555 = F_pg_cryptohash_init(m, v546)
	mBase = m.M
	if int32(0) <= v555 {
		v563 = v546
		goto L169
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	m.T0[v560].(func(*base.Module, int32, int32, int32))(m, v461, int32(83566), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v563 = v546
	goto L169
L178:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	m.T0[v569].(func(*base.Module, int32, int32, int32))(m, v461, int32(83608), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L6
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v575 = F_pg_cryptohash_final(m, v563, v459+int32(80), int32(32))
	mBase = m.M
	if v575 < int32(0) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	goto L180
L182:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	m.T0[v580].(func(*base.Module, int32, int32, int32))(m, v461, int32(83526), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L6
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	if v583 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	goto L184
L186:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	m.T0[v589].(func(*base.Module, int32, int32, int32))(m, v586, int32(299721), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L6
	} else {
		goto L189
	}
L187:
	;
	v593 = v583
	goto L188
L188:
	;
	v594 = F_strlen(m, v593)
	mBase = m.M
	if v594 != int32(64) {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v593 = v592
	goto L188
L190:
	;
	v730 = v459 + int32(80)
	v732 = v459 + int32(48)
	v733 = int32(32)
	goto L218
L191:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v459)+16)) = v593
	m.T0[v704].(func(*base.Module, int32, int32, int32))(m, v461, int32(759017), v459+int32(16))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L6
	} else {
		goto L214
	}
L192:
	;
	v607 = int32(0)
	goto L193
L193:
	;
	v618 = v593 + v607<<(uint(int32(1))%32)
	v619 = int32(*(*int8)(unsafe.Add(mBase, uint32(v618))))
	v621 = v619 - int32(48)
	if base.Ui32(v621&int32(255)) <= base.Ui32(int32(9)) {
		v644 = v621
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L190
L195:
	;
	v645 = int32(*(*int8)(unsafe.Add(mBase, uint32(v618)+1)))
	v647 = v645 - int32(48)
	if base.Ui32(v647&int32(255)) <= base.Ui32(int32(9)) {
		v670 = v647
		goto L203
	} else {
		goto L204
	}
L196:
	;
	if base.Ui32((v619-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v644 = v619 - int32(87)
	goto L195
L198:
	;
	goto L199
L199:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v619-int32(65))&int32(255)) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v643 = int32(-1)
	goto L202
L201:
	;
	v643 = v619 - int32(55)
	goto L202
L202:
	;
	v644 = v643
	goto L195
L203:
	;
	if v644 < int32(0) {
		goto L191
	} else {
		goto L211
	}
L204:
	;
	if base.Ui32((v645-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v670 = v645 - int32(87)
	goto L203
L206:
	;
	goto L207
L207:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v645-int32(65))&int32(255)) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v669 = int32(-1)
	goto L210
L209:
	;
	v669 = v645 - int32(55)
	goto L210
L210:
	;
	v670 = v669
	goto L203
L211:
	;
	if v670 < int32(0) {
		goto L191
	} else {
		goto L212
	}
L212:
	;
	v680 = v670 + v644<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v459+int32(48)+v607))) = uint8(v680)
	v683 = v607 + int32(1)
	if v683 != int32(32) {
		v607 = v683
		goto L193
	} else {
		goto L213
	}
L213:
	;
	goto L194
L214:
	;
	goto L190
L215:
	;
	if v795 != 0 {
		goto L233
	} else {
		goto L234
	}
L216:
	;
	v795 = int32(0)
	goto L215
L217:
	;
	v769 = v764
	v770 = v765
	v771 = v766
	goto L227
L218:
	;
	if (v730|v732)&int32(3) != 0 {
		v764 = v730
		v765 = v732
		v766 = v733
		goto L217
	} else {
		goto L221
	}
L220:
	;
	if v754 == int32(0) {
		goto L216
	} else {
		goto L226
	}
L221:
	;
	v741 = v730
	v742 = v732
	v743 = v733
	goto L222
L222:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v742)))
	if v746 != v747 {
		v764 = v741
		v765 = v742
		v766 = v743
		goto L217
	} else {
		goto L224
	}
L223:
	;
	goto L220
L224:
	;
	v749 = int32(4)
	v750 = v742 + v749
	v752 = v741 + v749
	v754 = v743 - v749
	if base.Ui32(int32(3)) < base.Ui32(v754) {
		v741 = v752
		v742 = v750
		v743 = v754
		goto L222
	} else {
		goto L225
	}
L225:
	;
	goto L223
L226:
	;
	v764 = v752
	v765 = v750
	v766 = v754
	goto L217
L227:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769))))
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
	if v774 == v775 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v795 = v774 - v775
	goto L215
L229:
	;
	v777 = int32(1)
	v782 = v771 - v777
	if v782 != 0 {
		v769 = v769 + v777
		v770 = v770 + v777
		v771 = v782
		goto L227
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	goto L228
L232:
	;
	goto L216
L233:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	m.T0[v798].(func(*base.Module, int32, int32, int32))(m, v461, int32(340045), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L6
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	F_pg_cryptohash_free(m, v563)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L6
	} else {
		goto L237
	}
L236:
	;
	goto L235
L237:
	;
	m.G0 = v459 + int32(112)
	goto L136
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v859
	m.T0[v861].(func(*base.Module, int32, int32, int32))(m, v24, int32(209701), v21+int32(16))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L6
	} else {
		goto L241
	}
L241:
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
	v6 = F_populate_record_worker(m, l0, int32(440222), v3, v3, int32(0))
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v3 = m.G0
	v5 = v3 - int32(80)
	m.G0 = v5
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v15 = F_pq_getmsgtext(m, v9, v10-v11, v5+int32(76))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+76))
		v21 = *(*int32)(unsafe.Add(mBase, _consts[462]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
		v24 = F_makeJsonLexContextCstringLen(m, v5+int32(8), v15, v19, v22, int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v30 = F_pg_parse_json_or_errsave(m, v5+int32(8), int32(1888256), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)+76))
				v33 = F_cstring_to_text_with_len(m, v15, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(80)
					return v33
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
	var v25 int32
	_ = v25
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v5 = l4
	if l0 == int32(0) {
		v9 = F_palloc0(m, int32(68))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				return int32(4555704)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v17 | int32(1)
				v25 = v9
				*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = l1
				*(*uint8)(unsafe.Add(mBase, uint32(v25)+56)) = uint8(v5)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l2
				if v5 != 0 {
					v36 = F_makeStringInfo(m)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v25)+60)) = v36
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v39 | int32(2)
						return v25
					}
				} else {
					return v25
				}
			}
		}
	} else {
		v24 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), int32(68))
		mBase = m.M
		v25 = l0
		*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v25))) = l1
		*(*uint8)(unsafe.Add(mBase, uint32(v25)+56)) = uint8(v5)
		*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l2
		if v5 != 0 {
			v36 = F_makeStringInfo(m)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v25)+60)) = v36
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v39 | int32(2)
				return v25
			}
		} else {
			return v25
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v731 {
		goto L296
	} else {
		goto L297
	}
L8:
	;
	F_appendStringInfoString(m, l0, int32(317374))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L295
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L292
	}
L10:
	;
	F_appendStringInfoString(m, l0, int32(715439))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L285
	}
L11:
	;
	F_appendStringInfoString(m, l0, int32(715527))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L278
	}
L12:
	;
	F_appendStringInfoString(m, l0, int32(715454))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L271
	}
L13:
	;
	F_appendStringInfoString(m, l0, int32(715569))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L264
	}
L14:
	;
	F_appendStringInfoString(m, l0, int32(714188))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L263
	}
L15:
	;
	F_appendStringInfoString(m, l0, int32(713891))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L262
	}
L16:
	;
	F_appendStringInfoString(m, l0, int32(713748))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L261
	}
L17:
	;
	F_appendStringInfoString(m, l0, int32(715548))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L248
	}
L18:
	;
	F_appendStringInfoString(m, l0, int32(714251))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L247
	}
L19:
	;
	F_appendStringInfoString(m, l0, int32(714082))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L246
	}
L20:
	;
	F_appendStringInfoString(m, l0, int32(713221))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
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
	F_appendStringInfoString(m, l0, int32(83989))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L201
	}
L23:
	;
	F_appendStringInfoString(m, l0, int32(714217))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L200
	}
L24:
	;
	F_appendStringInfoString(m, l0, int32(715558))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L193
	}
L25:
	;
	F_appendStringInfoString(m, l0, int32(714344))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L192
	}
L26:
	;
	F_appendStringInfoString(m, l0, int32(714198))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L191
	}
L27:
	;
	F_appendStringInfoString(m, l0, int32(713723))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L190
	}
L28:
	;
	F_appendStringInfoString(m, l0, int32(713507))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L189
	}
L29:
	;
	F_appendStringInfoString(m, l0, int32(714209))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L188
	}
L30:
	;
	F_appendStringInfoString(m, l0, int32(714295))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L187
	}
L31:
	;
	F_appendStringInfoString(m, l0, int32(715934))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L183
	}
L32:
	;
	F_appendStringInfoString(m, l0, int32(715895))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L179
	}
L33:
	;
	F_appendStringInfoChar(m, l0, int32(36))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L177
	}
L34:
	;
	F_appendStringInfoChar(m, l0, int32(36))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L176
	}
L35:
	;
	F_appendStringInfoChar(m, l0, int32(64))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
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
	v240 = m.ExcPending
	if v240 != 0 {
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
	F_appendStringInfoString(m, l0, int32(532400))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
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
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L92
	}
L43:
	;
	F_appendStringInfoString(m, l0, int32(715931))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
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
	F_appendStringInfoString(m, l0, int32(361081))
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
	F_appendStringInfoString(m, l0, int32(378600))
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
	v59 = int32(6)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v63 = v61 - int32(4)
	if base.Ui32(v63) <= base.Ui32(int32(37)) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63<<(uint(int32(2))%32))+uint32(_consts[988])))
	v71 = v70
	goto L69
L68:
	;
	v71 = v59
	goto L69
L69:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v77 = v75 - int32(4)
	if base.Ui32(v77) <= base.Ui32(int32(37)) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77<<(uint(int32(2))%32))+uint32(_consts[988])))
	v85 = v84
	goto L72
L71:
	;
	v85 = v59
	goto L72
L72:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v71) <= base.Ui32(v85)))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_appendStringInfoChar(m, l0, int32(32))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v93 = F_jspOperationName(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_appendStringInfoString(m, l0, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_appendStringInfoChar(m, l0, int32(32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_jspInitByBuffer(m, v10+int32(100), v102, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v109 = v107 - int32(4)
	if base.Ui32(v109) <= base.Ui32(int32(37)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v109<<(uint(int32(2))%32))+uint32(_consts[988])))
	v117 = v116
	goto L81
L80:
	;
	v117 = int32(6)
	goto L81
L81:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v123 = v121 - int32(4)
	if base.Ui32(v123) <= base.Ui32(int32(37)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v123<<(uint(int32(2))%32))+uint32(_consts[988])))
	v131 = v130
	goto L84
L83:
	;
	v131 = v59
	goto L84
L84:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v117) <= base.Ui32(v131)))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
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
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L7
L88:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v145, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v151 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v151, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	goto L7
L92:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v163, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v169 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v169, v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_appendStringInfoString(m, l0, int32(255659))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
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
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	v182 = v18
	goto L98
L98:
	;
	if v182 == int32(19) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v182 = v181
	goto L98
L100:
	;
	v185 = int32(43)
	goto L102
L101:
	;
	v185 = int32(45)
	goto L102
L102:
	;
	F_appendStringInfoChar(m, l0, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v190, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v194 = int32(6)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v198 = v196 - int32(4)
	if base.Ui32(v198) <= base.Ui32(int32(37)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v198<<(uint(int32(2))%32))+uint32(_consts[988])))
	v206 = v205
	goto L107
L106:
	;
	v206 = v194
	goto L107
L107:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v212 = v210 - int32(4)
	if base.Ui32(v212) <= base.Ui32(int32(37)) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v212<<(uint(int32(2))%32))+uint32(_consts[988])))
	v220 = v219
	goto L110
L109:
	;
	v220 = v194
	goto L110
L110:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v206) <= base.Ui32(v220)))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
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
	v228 = m.ExcPending
	if v228 != 0 {
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
	v234 = m.ExcPending
	if v234 != 0 {
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
	v237 = m.ExcPending
	if v237 != 0 {
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
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v241 <= int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L146
	}
L122:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	F_jspInitByBuffer(m, v10+int32(100), v246, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v252 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v10+int32(72), v255, v252)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	v271 = v10 + int32(100)
	goto L126
L126:
	;
	v272 = int32(0)
	F_printJsonPathItem(m, l0, v271, v272, v272)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L130
	}
L127:
	;
	v260 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v260, v260)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_appendStringInfoString(m, l0, int32(771912))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v271 = v10 + int32(72)
	goto L126
L130:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v276 < int32(2) {
		goto L121
	} else {
		goto L131
	}
L131:
	;
	v283 = int32(1)
	goto L132
L132:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v291 = v283 << (uint(int32(3)) % 32)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v291+v292)))
	F_jspInitByBuffer(m, v10+int32(100), v289, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L134
	}
L133:
	;
	goto L121
L134:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v297+v291)+4))
	if v299 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v325 = int32(0)
	F_printJsonPathItem(m, l0, v324, v325, v325)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L144
	}
L136:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v10+int32(72), v302, v299)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
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
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L143
	}
L139:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v310 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v310, v310)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_appendStringInfoString(m, l0, int32(771912))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v324 = v10 + int32(72)
	goto L135
L143:
	;
	v324 = v10 + int32(100)
	goto L135
L144:
	;
	v330 = v283 + int32(1)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v330 < v331 {
		v283 = v330
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
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v347 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	goto L149
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v347
	F_appendStringInfo(m, l0, int32(4249), v10+int32(16))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L169
	}
L152:
	;
	if v347 == int32(-1) {
		goto L163
	} else {
		goto L164
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v347
	F_appendStringInfo(m, l0, int32(4242), v10+int32(32))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L162
	}
L154:
	;
	switch v346 + int32(1) {
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
	if v346 != v347 {
		goto L152
	} else {
		goto L159
	}
L157:
	;
	F_appendStringInfoString(m, l0, int32(699109))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	goto L7
L159:
	;
	if v347 != int32(-1) {
		goto L153
	} else {
		goto L160
	}
L160:
	;
	F_appendStringInfoString(m, l0, int32(4322))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v346
	F_appendStringInfo(m, l0, int32(4262), v10+int32(48))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	if v346 != int32(-1) {
		goto L151
	} else {
		goto L167
	}
L166:
	;
	goto L7
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v347
	F_appendStringInfo(m, l0, int32(4331), v10-int32(-64))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
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
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_escape_json_with_len(m, l0, v393, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
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
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_escape_json_with_len(m, l0, v406, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L7
L179:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v415, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v421 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v421, v421)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	goto L7
L183:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v433, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v439 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v439, v439)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
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
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v467 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v10+int32(100), v470, v467)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
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
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L199
	}
L197:
	;
	v475 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v475, v475)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
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
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v493, v494)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L206
	}
L205:
	;
	goto L204
L206:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v504 = v500 - int32(4)
	if base.Ui32(v504) <= base.Ui32(int32(37)) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v517 = v513 - int32(4)
	if base.Ui32(v517) <= base.Ui32(int32(37)) {
		goto L212
	} else {
		goto L213
	}
L208:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v504<<(uint(int32(2))%32))+uint32(_consts[988])))
	v512 = v511
	goto L210
L209:
	;
	v512 = int32(6)
	goto L210
L210:
	;
	goto L207
L211:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v512) <= base.Ui32(v525)))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L215
	}
L212:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v517<<(uint(int32(2))%32))+uint32(_consts[988])))
	v525 = v524
	goto L214
L213:
	;
	v525 = int32(6)
	goto L214
L214:
	;
	goto L211
L215:
	;
	F_appendStringInfoString(m, l0, int32(764004))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_escape_json_with_len(m, l0, v532, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v536 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	F_appendStringInfoString(m, l0, int32(763580))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
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
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v540&int32(1) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	F_appendStringInfoChar(m, l0, int32(105))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	v547 = v540
	goto L224
L224:
	;
	if v547&int32(2) != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v547 = v546
	goto L224
L226:
	;
	F_appendStringInfoChar(m, l0, int32(115))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	v554 = v547
	goto L228
L228:
	;
	if v554&int32(4) != 0 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v554 = v553
	goto L228
L230:
	;
	F_appendStringInfoChar(m, l0, int32(109))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	v561 = v554
	goto L232
L232:
	;
	if v561&int32(8) != 0 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v561 = v560
	goto L232
L234:
	;
	F_appendStringInfoChar(m, l0, int32(120))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	v568 = v561
	goto L236
L236:
	;
	if v568&int32(16) != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v568 = v567
	goto L236
L238:
	;
	F_appendStringInfoChar(m, l0, int32(113))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
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
	v576 = m.ExcPending
	if v576 != 0 {
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
	v582 = m.ExcPending
	if v582 != 0 {
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
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v595 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v10+int32(100), v598, v595)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v607 != 0 {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	v603 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v603, v603)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
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
	v610 = m.ExcPending
	if v610 != 0 {
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
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L260
	}
L257:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_jspInitByBuffer(m, v10+int32(100), v613, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v619 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v619, v619)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
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
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v638 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v10+int32(100), v641, v638)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
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
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	v646 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v646, v646)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
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
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v656 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v10+int32(100), v659, v656)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
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
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	v664 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v664, v664)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
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
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v674 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v10+int32(100), v677, v674)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
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
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L284
	}
L282:
	;
	v682 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v682, v682)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
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
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v692 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v10+int32(100), v695, v692)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
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
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	v700 = int32(0)
	F_printJsonPathItem(m, l0, v10+int32(100), v700, v700)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
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
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v711
	F_errmsg_internal(m, int32(506940), v10)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(521387), int32(835), int32(304859))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
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
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v10+int32(100), v736, v731)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
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
	v741 = int32(1)
	F_printJsonPathItem(m, l0, v10+int32(100), v741, v741)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	goto L298
}
