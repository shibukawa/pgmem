package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_seg_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 float32
	_ = v18
	var v19 int32
	_ = v19
	var v20 float32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 float32
	_ = v75
	var v76 float32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = int32(-1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*float32)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*float32)(unsafe.Add(mBase, uint32(v19)))
	if base.F32_lt(v18, v20) != 0 {
		v138 = v16
		m.G0 = v14 + int32(32)
		return v138
	} else {
		if base.F32_gt(v18, v20) != 0 {
			v138 = int32(1)
			m.G0 = v14 + int32(32)
			return v138
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+10)))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+10)))
			v25 = base.B2i32(v23 == v24)
			if v25 == int32(0) {
				if v23 == int32(45) {
					v138 = v16
					m.G0 = v14 + int32(32)
					return v138
				} else {
					if v24 == int32(45) {
						v138 = int32(1)
						m.G0 = v14 + int32(32)
						return v138
					} else {
						if v23 == int32(60) {
							v138 = v16
							m.G0 = v14 + int32(32)
							return v138
						} else {
							if v23 == int32(62) {
								v138 = int32(1)
								m.G0 = v14 + int32(32)
								return v138
							} else {
								if v24 == int32(60) {
									v138 = int32(1)
									m.G0 = v14 + int32(32)
									return v138
								} else {
									if v24 == int32(62) {
										v138 = v16
										m.G0 = v14 + int32(32)
										return v138
									} else {
										v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+8)))
										v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+8)))
										if v40 < v41 {
											v138 = v16
											m.G0 = v14 + int32(32)
											return v138
										} else {
											if v41 < v40 {
												v138 = int32(1)
												m.G0 = v14 + int32(32)
												return v138
											} else {
												if v25 == int32(0) {
													if v23 == int32(126) {
														v138 = v16
														m.G0 = v14 + int32(32)
														return v138
													} else {
														if v24 == int32(126) {
															v138 = int32(1)
															m.G0 = v14 + int32(32)
															return v138
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v56 = m.ExcPending
															if v56 != 0 {
																return int32(0)
															} else {
																v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+10)))
																v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+10)))
																*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v58
																*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v57
																F_errmsg_internal(m, int32(501001), v14+int32(16))
																mBase = m.M
																v67 = m.ExcPending
																if v67 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(521836), int32(787), int32(247493))
																	mBase = m.M
																	v74 = m.ExcPending
																	if v74 != 0 {
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
													v75 = *(*float32)(unsafe.Add(mBase, uint32(v17)+4))
													v76 = *(*float32)(unsafe.Add(mBase, uint32(v19)+4))
													if base.F32_lt(v75, v76) != 0 {
														v138 = v16
														m.G0 = v14 + int32(32)
														return v138
													} else {
														if base.F32_gt(v75, v76) != 0 {
															v138 = int32(1)
															m.G0 = v14 + int32(32)
															return v138
														} else {
															v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+11)))
															v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
															v81 = base.B2i32(v79 == v80)
															if v81 == int32(0) {
																if v79 == int32(45) {
																	v138 = int32(1)
																	m.G0 = v14 + int32(32)
																	return v138
																} else {
																	if v79 == int32(60) {
																		v138 = v16
																		m.G0 = v14 + int32(32)
																		return v138
																	} else {
																		if v80 == int32(45) {
																			v138 = v16
																			m.G0 = v14 + int32(32)
																			return v138
																		} else {
																			if v79 == int32(62) {
																				v138 = int32(1)
																				m.G0 = v14 + int32(32)
																				return v138
																			} else {
																				if v80 == int32(60) {
																					v138 = int32(1)
																					m.G0 = v14 + int32(32)
																					return v138
																				} else {
																					if v80 == int32(62) {
																						v138 = v16
																						m.G0 = v14 + int32(32)
																						return v138
																					} else {
																						v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+9)))
																						v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+9)))
																						if v96 < v97 {
																							v138 = int32(1)
																							m.G0 = v14 + int32(32)
																							return v138
																						} else {
																							v101 = base.B2i32(v97 < v96)
																							if v97 < v96 {
																								v102 = int32(-1)
																							} else {
																								v102 = int32(0)
																							}
																							if v79 == v80 {
																								v138 = v102
																								m.G0 = v14 + int32(32)
																								return v138
																							} else {
																								if v97 < v96 {
																									v138 = v102
																									m.G0 = v14 + int32(32)
																									return v138
																								} else {
																									if v79 == int32(126) {
																										v138 = int32(1)
																										m.G0 = v14 + int32(32)
																										return v138
																									} else {
																										if v80 == int32(126) {
																											v138 = int32(-1)
																											m.G0 = v14 + int32(32)
																											return v138
																										} else {
																											F_errstart_cold(m, int32(21), int32(0))
																											mBase = m.M
																											v112 = m.ExcPending
																											if v112 != 0 {
																												return int32(0)
																											} else {
																												v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+11)))
																												v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+11)))
																												*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v114
																												*(*int32)(unsafe.Add(mBase, uint32(v14))) = v113
																												F_errmsg_internal(m, int32(501034), v14)
																												mBase = m.M
																												v121 = m.ExcPending
																												if v121 != 0 {
																													return int32(0)
																												} else {
																													F_errfinish(m, int32(521836), int32(845), int32(247493))
																													mBase = m.M
																													v128 = m.ExcPending
																													if v128 != 0 {
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
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+9)))
																v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+9)))
																if v96 < v97 {
																	v138 = int32(1)
																	m.G0 = v14 + int32(32)
																	return v138
																} else {
																	v101 = base.B2i32(v97 < v96)
																	if v97 < v96 {
																		v102 = int32(-1)
																	} else {
																		v102 = int32(0)
																	}
																	if v79 == v80 {
																		v138 = v102
																		m.G0 = v14 + int32(32)
																		return v138
																	} else {
																		if v97 < v96 {
																			v138 = v102
																			m.G0 = v14 + int32(32)
																			return v138
																		} else {
																			if v79 == int32(126) {
																				v138 = int32(1)
																				m.G0 = v14 + int32(32)
																				return v138
																			} else {
																				if v80 == int32(126) {
																					v138 = int32(-1)
																					m.G0 = v14 + int32(32)
																					return v138
																				} else {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v112 = m.ExcPending
																					if v112 != 0 {
																						return int32(0)
																					} else {
																						v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+11)))
																						v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+11)))
																						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v114
																						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v113
																						F_errmsg_internal(m, int32(501034), v14)
																						mBase = m.M
																						v121 = m.ExcPending
																						if v121 != 0 {
																							return int32(0)
																						} else {
																							F_errfinish(m, int32(521836), int32(845), int32(247493))
																							mBase = m.M
																							v128 = m.ExcPending
																							if v128 != 0 {
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
				}
			} else {
				v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+8)))
				v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+8)))
				if v40 < v41 {
					v138 = v16
					m.G0 = v14 + int32(32)
					return v138
				} else {
					if v41 < v40 {
						v138 = int32(1)
						m.G0 = v14 + int32(32)
						return v138
					} else {
						if v25 == int32(0) {
							if v23 == int32(126) {
								v138 = v16
								m.G0 = v14 + int32(32)
								return v138
							} else {
								if v24 == int32(126) {
									v138 = int32(1)
									m.G0 = v14 + int32(32)
									return v138
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+10)))
										v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+10)))
										*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v58
										*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v57
										F_errmsg_internal(m, int32(501001), v14+int32(16))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(521836), int32(787), int32(247493))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
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
							v75 = *(*float32)(unsafe.Add(mBase, uint32(v17)+4))
							v76 = *(*float32)(unsafe.Add(mBase, uint32(v19)+4))
							if base.F32_lt(v75, v76) != 0 {
								v138 = v16
								m.G0 = v14 + int32(32)
								return v138
							} else {
								if base.F32_gt(v75, v76) != 0 {
									v138 = int32(1)
									m.G0 = v14 + int32(32)
									return v138
								} else {
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+11)))
									v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
									v81 = base.B2i32(v79 == v80)
									if v81 == int32(0) {
										if v79 == int32(45) {
											v138 = int32(1)
											m.G0 = v14 + int32(32)
											return v138
										} else {
											if v79 == int32(60) {
												v138 = v16
												m.G0 = v14 + int32(32)
												return v138
											} else {
												if v80 == int32(45) {
													v138 = v16
													m.G0 = v14 + int32(32)
													return v138
												} else {
													if v79 == int32(62) {
														v138 = int32(1)
														m.G0 = v14 + int32(32)
														return v138
													} else {
														if v80 == int32(60) {
															v138 = int32(1)
															m.G0 = v14 + int32(32)
															return v138
														} else {
															if v80 == int32(62) {
																v138 = v16
																m.G0 = v14 + int32(32)
																return v138
															} else {
																v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+9)))
																v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+9)))
																if v96 < v97 {
																	v138 = int32(1)
																	m.G0 = v14 + int32(32)
																	return v138
																} else {
																	v101 = base.B2i32(v97 < v96)
																	if v97 < v96 {
																		v102 = int32(-1)
																	} else {
																		v102 = int32(0)
																	}
																	if v79 == v80 {
																		v138 = v102
																		m.G0 = v14 + int32(32)
																		return v138
																	} else {
																		if v97 < v96 {
																			v138 = v102
																			m.G0 = v14 + int32(32)
																			return v138
																		} else {
																			if v79 == int32(126) {
																				v138 = int32(1)
																				m.G0 = v14 + int32(32)
																				return v138
																			} else {
																				if v80 == int32(126) {
																					v138 = int32(-1)
																					m.G0 = v14 + int32(32)
																					return v138
																				} else {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v112 = m.ExcPending
																					if v112 != 0 {
																						return int32(0)
																					} else {
																						v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+11)))
																						v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+11)))
																						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v114
																						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v113
																						F_errmsg_internal(m, int32(501034), v14)
																						mBase = m.M
																						v121 = m.ExcPending
																						if v121 != 0 {
																							return int32(0)
																						} else {
																							F_errfinish(m, int32(521836), int32(845), int32(247493))
																							mBase = m.M
																							v128 = m.ExcPending
																							if v128 != 0 {
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
																}
															}
														}
													}
												}
											}
										}
									} else {
										v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+9)))
										v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+9)))
										if v96 < v97 {
											v138 = int32(1)
											m.G0 = v14 + int32(32)
											return v138
										} else {
											v101 = base.B2i32(v97 < v96)
											if v97 < v96 {
												v102 = int32(-1)
											} else {
												v102 = int32(0)
											}
											if v79 == v80 {
												v138 = v102
												m.G0 = v14 + int32(32)
												return v138
											} else {
												if v97 < v96 {
													v138 = v102
													m.G0 = v14 + int32(32)
													return v138
												} else {
													if v79 == int32(126) {
														v138 = int32(1)
														m.G0 = v14 + int32(32)
														return v138
													} else {
														if v80 == int32(126) {
															v138 = int32(-1)
															m.G0 = v14 + int32(32)
															return v138
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return int32(0)
															} else {
																v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+11)))
																v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+11)))
																*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v114
																*(*int32)(unsafe.Add(mBase, uint32(v14))) = v113
																F_errmsg_internal(m, int32(501034), v14)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(521836), int32(845), int32(247493))
																	mBase = m.M
																	v128 = m.ExcPending
																	if v128 != 0 {
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
func F_seg_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(6670), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v6^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_seg_inter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 float32
	_ = v16
	var v17 float32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float32
	_ = v24
	var v26 float32
	_ = v26
	var v27 float32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 float32
	_ = v34
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_palloc(m, int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*float32)(unsafe.Add(mBase, uint32(v9)+4))
		v17 = *(*float32)(unsafe.Add(mBase, uint32(v10)+4))
		v18 = base.F32_lt(v16, v17)
		if v18 != 0 {
			v19 = v9
		} else {
			v19 = v10
		}
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+9)))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v21)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)) = uint8(v20)
		if v18 != 0 {
			v24 = v16
		} else {
			v24 = v17
		}
		*(*float32)(unsafe.Add(mBase, uint32(v12)+4)) = v24
		v26 = *(*float32)(unsafe.Add(mBase, uint32(v9)))
		v27 = *(*float32)(unsafe.Add(mBase, uint32(v10)))
		v28 = base.F32_gt(v26, v27)
		if v28 != 0 {
			v29 = v9
		} else {
			v29 = v10
		}
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+10)))
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v31)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v30)
		if v28 != 0 {
			v34 = v26
		} else {
			v34 = v27
		}
		*(*float32)(unsafe.Add(mBase, uint32(v12))) = v34
		return v12
	}
}
func F_seg_yy_create_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v8 = F_palloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
			v15 = F_palloc(m, l1+int32(2))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_7(m, int32(713847))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v20
					v23 = *(*int32)(unsafe.Add(mBase, _consts[163]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v24)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v20
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					if v37 == v24 {
						v59 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v59
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v59
						v66 = int32(0)
						v67 = int32(36)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v37+v40<<(uint(int32(2))%32))))
						if v8 != v44 {
							v59 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v59
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v59
							v66 = int32(0)
							v67 = int32(36)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v48
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v51
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)) = uint8(v53)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							v66 = int32(1)
							v67 = int32(40)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v67+v8))) = v66
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[163])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_7(m, int32(713847))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_seg_yy_scan_bytes(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	v4 = int32(0)
	if base.Ui32(l1) < base.Ui32(int32(-2)) {
		v14 = l1 + int32(2)
		v15 = F_palloc(m, v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				F_yy_fatal_error_7(m, int32(713374))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if l1 == int32(0) {
				} else {
					if base.Ui32(int32(4)) <= base.Ui32(l1) {
						v30 = v4
						v34 = v4
						for {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v30))))
							*(*uint8)(unsafe.Add(mBase, uint32(v30+v15))) = uint8(v39)
							v42 = v30 | int32(1)
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v42))))
							*(*uint8)(unsafe.Add(mBase, uint32(v15+v42))) = uint8(v45)
							v48 = v30 | int32(2)
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v48))))
							*(*uint8)(unsafe.Add(mBase, uint32(v15+v48))) = uint8(v51)
							v54 = v30 | int32(3)
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v54))))
							*(*uint8)(unsafe.Add(mBase, uint32(v15+v54))) = uint8(v57)
							v59 = int32(4)
							v60 = v30 + v59
							v62 = v34 + v59
							if v62 != l1&int32(-4) {
								v30 = v60
								v34 = v62
								continue
							} else {
								break
							}
							break
						}
						v67 = v60
					} else {
						v67 = v4
					}
					v75 = l1 & int32(3)
					if v75 == int32(0) {
					} else {
						v81 = v67
						v84 = v4
						for {
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v81))))
							*(*uint8)(unsafe.Add(mBase, uint32(v81+v15))) = uint8(v90)
							v92 = int32(1)
							v95 = v84 + v92
							if v95 != v75 {
								v81 = v81 + v92
								v84 = v95
								continue
							} else {
								break
							}
							break
						}
					}
				}
				v108 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1+v15))) = uint16(v108)
				v110 = F_seg_yy_scan_buffer(m, v15, v14, l2)
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					if v110 == int32(0) {
						F_yy_fatal_error_7(m, int32(713415))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = int32(1)
						return v110
					}
				}
			}
		}
	} else {
		F_yy_fatal_error_7(m, int32(713445))
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
